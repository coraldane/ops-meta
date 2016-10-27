package user

import (
	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/cache"
	"github.com/coraldane/ops-meta/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Invalidate() {
	this.DelSession(cache.SESSION_UID)
	this.Ctx.Redirect(302, "/user/login")
}

func (this *UserController) NoPermission() {
	this.Layout = "layout/default_layout.html"
	this.TplName = "user/no_permission.html"
}

func (this *UserController) Query() {
	jsonResult := make(map[string]interface{})

	pageInfo := models.NewPageInfo(this.Ctx.Request)

	queryDto := models.QueryUserDto{}
	if err := this.ParseForm(&queryDto); nil != err {
		logger.Errorln("parseForm error", err)
	}

	list, pageInfo := models.QueryUserList(queryDto, pageInfo)
	jsonResult["items"] = list
	jsonResult["total"] = pageInfo.RowCount

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}
