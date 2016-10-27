package api

import (
	"strings"

	"github.com/astaxie/beego"
	// "github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) Reload() {
	jsonResult := make(map[string]interface{})

	if strings.HasPrefix(this.Ctx.Request.RemoteAddr, "127.0.0.1") {
		err := g.ParseConfig(g.ConfigFile)

		if nil == err {
			this.Data["json"] = g.Config()
		} else {
			jsonResult["message"] = err.Error()
			this.Data["json"] = &jsonResult
		}
		this.ServeJSON()
	} else {
		this.Ctx.WriteString("no privilege")
	}
}
