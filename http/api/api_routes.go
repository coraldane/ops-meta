package api

import (
	"github.com/astaxie/beego"
)

func ConfigRoutes() {
	beego.Router("/api/config/reload", &ConfigController{}, "get:Reload")
	beego.Router("/api/health/check", &HealthController{}, "get:Check")

	beego.Router("/api/version", &FrameController{}, "get:Version")
	beego.Router("/api/workdir", &FrameController{}, "get:WorkDir")

	beego.Router("/api/heartbeat", &HeartbeatController{}, "post:Post")
}
