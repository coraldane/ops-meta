package api

import (
	"github.com/astaxie/beego"
)

type HealthController struct {
	beego.Controller
}

func (this *HealthController) Check() {
	jsonResult := make(map[string]interface{})

	jsonResult["success"] = true
	jsonResult["message"] = "ok"
	this.Data["json"] = &jsonResult

	this.ServeJSON()
}
