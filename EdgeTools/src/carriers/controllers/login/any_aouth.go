package login

import (
	"carriers/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

/*
 * 统一登录验证地址
 * */

type AnyOauth struct {
	beego.Controller
}

func (c *AnyOauth) Get() {
	beego.Error("LoginOauth use get method!!! Host:", c.Ctx.Request.URL.Host)
	c.Ctx.WriteString("{}")
}

func (c *AnyOauth) Post() {
	beego.Info("AnyLogin app host :", c.Ctx.Request.URL.Host)
	beego.Debug("AnyLoginOauth Post body length :", len(c.Ctx.Input.RequestBody))
	beego.Debug("AnyLoginOauth Post body content :", string(c.Ctx.Input.RequestBody))

	req := httplib.Post(beego.AppConfig.String("anyoauthurl"))
	req.Header("content-type", "application/x-www-form-urlencoded")
	req.Body(c.Ctx.Input.RequestBody)

	str, err := req.String()

	if err != nil {
		// error
		c.Ctx.WriteString("{\"status\":\"fail\"}")
		return
	}
	beego.Debug("LoginOauth Req oauth server result : ", str)

	var oauthObj models.OauthObject
	json.Unmarshal([]byte(str), &oauthObj)

	b, _ := beego.AppConfig.Bool("Debug")
	if b {
		c.Ctx.WriteString(str)
		return
	}

	beego.Debug("Login Status:", oauthObj.Status)
	if oauthObj.Status == "ok" {

		servId, _ := strconv.Atoi(c.Input().Get("server_id"))
		host := models.GetServOauthHostById(servId)
		gamereq := httplib.Post(host)
		gamereq.Body([]byte(str))

		_, gameerr := gamereq.String()
		if gameerr != nil {
			beego.Debug(gameerr.Error())
			c.Ctx.WriteString("{\"status\":\"fail\"}")
			return
		}

		c.Ctx.WriteString(str)
		return

	} else if oauthObj.Status == "fail" {
		c.Ctx.WriteString(str)
		return
	} else {
		c.Ctx.WriteString("{\"status\":\"fail\"}")
		return
	}
}
