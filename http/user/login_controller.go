package user

import (
	"github.com/astaxie/beego"
	"github.com/coraldane/ops-meta/cache"
	"github.com/coraldane/ops-meta/models"
	"github.com/toolkits/logger"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["originURL"] = this.GetSession("originURL")
	this.TplName = "user/login.html"
}

func (this *LoginController) Post() {
	jsonResult := make(map[string]interface{})
	jsonResult["success"] = false
	userName := this.GetString("userName")
	loginPwd := this.GetString("loginPwd")

	if "" == userName {
		jsonResult["message"] = "用户名不能为空"
	} else if "" == loginPwd {
		jsonResult["message"] = "登陆密码不能为空"
	} else {
		user, err := models.CheckLogin(userName, loginPwd)
		if nil != err {
			if "<QuerySeter> no row found" == err.Error() {
				jsonResult["message"] = "用户名或密码错误"
			} else {
				logger.Errorln("checkLogin error", userName, loginPwd, err)
				jsonResult["message"] = err.Error()
			}
		} else if nil == user {
			jsonResult["message"] = "用户名或密码错误"
		} else {
			this.SetSession(cache.SESSION_UID, user.Id)
			this.SetSession(cache.SESSION_USERNAME, user.RealName)
			this.SetSession(cache.SESSION_ROLENAME, user.RoleName)
			jsonResult["success"] = true

			originURL := this.GetString("originURL")
			if "" == originURL {
				originURL = ""
			}
			jsonResult["message"] = originURL
		}
	}

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}
