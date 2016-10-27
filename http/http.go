package http

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/cache"
	"github.com/coraldane/ops-meta/http/admin"
	"github.com/coraldane/ops-meta/http/api"
	"github.com/coraldane/ops-meta/http/home"
	"github.com/coraldane/ops-meta/http/user"
	"github.com/coraldane/ops-meta/models"
)

func Start() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterCheckLogin)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterSetBasepath)

	api.ConfigRoutes()
	home.ConfigRoutes()
	admin.ConfigRoutes()
	user.ConfigRoutes()

	beego.AddFuncMap("datetime", func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	})
	beego.AddFuncMap("empty", func(obj interface{}) bool {
		if _, ok := obj.(string); ok {
			return "" == obj
		} else if value, ok := obj.(int64); ok {
			return 0 == value
		} else {
			return nil == obj
		}
	})

	beego.Run()
}

func FilterSetBasepath(ctx *context.Context) {
	ctx.Output.Header("Cache-Control", "no-cache")
	ctx.Output.Header("Expires", "0")

	ctx.Input.SetData("CurrentActionPath", ctx.Input.URL())
	host := ctx.Input.Header("Host")
	if "" == host {
		host = ctx.Request.Host
	}
	ctx.Input.SetData("BasePath", fmt.Sprintf("%s://%s/", ctx.Input.Scheme(), host))

	uid := ctx.Input.Session(cache.SESSION_UID)
	if nil != uid {
		ctx.Input.SetData(cache.SESSION_UID, uid.(int64))
		ctx.Input.SetData(cache.SESSION_USERNAME, ctx.Input.Session(cache.SESSION_USERNAME))
		ctx.Input.SetData(cache.SESSION_ROLENAME, ctx.Input.Session(cache.SESSION_ROLENAME))
	}
}

func FilterCheckLogin(ctx *context.Context) {
	url := ctx.Request.RequestURI
	if strings.HasPrefix(url, "/user") || "" == url || strings.HasPrefix(url, "/api") {
		return
	}

	userId, ok := ctx.Input.Session(cache.SESSION_UID).(int64)
	user := models.GetUserById(userId)
	if !ok || nil == user {
		ctx.Output.Session("originURL", ctx.Request.RequestURI)
		ctx.Redirect(302, "/user/login")
	} else if strings.HasPrefix(url, "/root") && "ADMIN" != user.RoleName {
		ctx.Redirect(302, "/user/nopermission")
	}
}
