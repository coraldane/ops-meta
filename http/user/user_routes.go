package user

import (
	"github.com/astaxie/beego"
)

func ConfigRoutes() {
	beego.Router("/user/login", &LoginController{}, "get:Get")
	beego.Router("/user/login", &LoginController{}, "post:Post")

	beego.Router("/user/invalidate", &UserController{}, "*:Invalidate")
	beego.Router("/user/nopermission", &UserController{}, "get:NoPermission")

	beego.Router("/i/user/query", &UserController{}, "post:Query")
}
