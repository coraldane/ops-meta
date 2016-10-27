package home

import (
	"github.com/astaxie/beego"
)

func ConfigRoutes() {
	beego.Router("/", &HomeController{})
}
