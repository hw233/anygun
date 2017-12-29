package web

import (
	"carriers/models"
	"fmt"

	"encoding/json"

	"github.com/astaxie/beego"
)

type (
	ServerList struct {
		StatusCode int               `json:"statusCode"`
		ServerList map[string]string `json:"serverList"`
	}

	GetServerList struct {
		beego.Controller
	}
)

func (c *GetServerList) Get() {
	servs := models.GetServers()

	sl := ServerList{200, map[string]string{}}

	for _, sv := range servs {
		if sv.Id < 2000 {
			continue
		}
		sl.ServerList[fmt.Sprintf("%d", sv.Id)] = sv.Name
	}

	ret, err := json.Marshal(sl)

	if err != nil {
		beego.Error(err.Error())
		c.Ctx.WriteString("{\"statusCode\":300}")
	}

	c.Ctx.WriteString(string(ret))
}
