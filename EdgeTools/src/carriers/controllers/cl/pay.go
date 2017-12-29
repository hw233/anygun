package cl

import (
	"carriers/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
	"net/url"
	"bytes"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"strconv"
)

func instertPayOrder(gameOrderId, startTime string) {
	db, err := models.CarriersDB()

	if err != nil {
		beego.Error(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `CL_Order`(`GameOrderId`,`SDKOrderId`,`StartTime`,`NodifyTime`,`Payment`,`ShopId`,`Status`)VALUES(?,?,?,?,?,?,?)", gameOrderId, "", startTime, "", 0, 0, 0)

	if err != nil {
		beego.Error(err.Error())
		return
	}
}

func updatePayOrder(gameOrderId, sdkOrderId, notifyTime, shopId string, Payment float32, status int) {
	db, err := models.CarriersDB()

	if err != nil {
		beego.Error(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `CL_Order` SET `SDKOrderId`= ?,`NodifyTime`= ?,`Payment`= ?,`ShopId`= ?,`Status`= ? WHERE `GameOrderId` = ?", sdkOrderId, notifyTime, Payment, shopId, status, gameOrderId)

	if err != nil {
		beego.Error(err.Error())
		return
	}
}

type PayOrder struct {
	beego.Controller
}

func (c *PayOrder) Post() {
	uid := c.GetString("uid")
	now := time.Now()
	linkstring := uid + now.String()
	md5bytes := md5.Sum([]byte(linkstring))
	gameOrderId := hex.EncodeToString(md5bytes[:])
	gameOrderId = fmt.Sprintf("CL%s", strings.ToUpper(gameOrderId))
	c.Ctx.WriteString(gameOrderId)
	beego.Info(fmt.Sprintf("query order id [%s|%s]", uid, gameOrderId))

	go instertPayOrder(gameOrderId, now.Format("2006-01-02 15:04:05"))
}

type (
	Pay struct {
		beego.Controller
	}
)

type URIValues map[string]string

func (v URIValues) Encode(keys []string) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer

	for _, k := range keys {
		vs := v[k]
		prefix := url.QueryEscape(k) + "="

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(vs))

	}
	return buf.String()
}

var signKeyOrder = []string{
	"n_time",
	"appid",
	"o_id",
	"t_fee",
	"g_name",
	"g_body",
	"t_status",
}

func (c *Pay) PostServer(orderId , tm , shopid , amount  ,serverId , userId , roleId string){
	uri := url.Values{}

	uri.Set("order_id",orderId)
	uri.Set("order_type",fmt.Sprint(channel))
	uri.Set("pay_time",tm)
	uri.Set("product_count","1")
	uri.Set("product_id",shopid)
	uri.Set("game_user_id",roleId)
	uri.Set("pay_status","1")
	uri.Set("user_id",userId)
	uri.Set("amount",amount)

	b, _ := json.Marshal(uri)

	models.InsertOrderLog(string(b))
	servId ,_ := strconv.Atoi(serverId)
	host := models.GetServPayHostById(servId)
	gamereq := httplib.Post(host)
	gamereq.Body(b)

	_, gameerr := gamereq.String()
	if gameerr != nil {
		beego.Error(gameerr.Error())
		c.Ctx.WriteString("FAIL")
		return
	}
	beego.Debug(string(b))
	c.Ctx.WriteString("success")

}

func (c *Pay) Post() {

	beego.Debug(c.Ctx.Request.URL.RawQuery)

	appid := c.GetString("appid")
	n_time := c.GetString("n_time")
	o_id := c.GetString("o_id")
	t_fee := c.GetString("t_fee")
	g_name := c.GetString("g_name")
	g_body := c.GetString("g_body")
	t_status := c.GetString("t_status")
	o_sign := c.GetString("o_sign")
	u_param := c.GetString("u_param")
	o_orderid := c.GetString("o_orderid")

	beego.Debug(n_time,appid ,o_id ,t_fee ,g_name ,g_body ,t_status ,o_sign ,u_param ,o_orderid)

	uri := URIValues{}
	uri["n_time"]= fmt.Sprint(n_time)
	uri["appid"]=fmt.Sprint(appid)
	uri["o_id"]=fmt.Sprint(o_id)
	uri["t_fee"]=fmt.Sprint(t_fee)
	uri["g_name"]=fmt.Sprint(g_name)
	uri["g_body"]=fmt.Sprint(g_body)
	uri["t_status"]=fmt.Sprint(t_status)

	uriString := fmt.Sprintf("%s%s",uri.Encode(signKeyOrder),appkey)

	beego.Debug(uriString)

	signMd5String :=  md5.Sum([]byte(uriString))

	sign := hex.EncodeToString(signMd5String[:])

	if sign != o_sign {
		beego.Error(sign ,"!=", o_sign)
	}

	params := strings.Split(u_param,";")
	if len(params) == 4 {
		c.PostServer(o_id, n_time, params[3],t_fee,params[0],params[1],params[2])
		return
	}
	c.Ctx.WriteString("FAIL")

}
