package api

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-common/model"
	"github.com/coraldane/ops-meta/store"
)

type HeartbeatController struct {
	beego.Controller
}

func (this *HeartbeatController) Post() {
	jsonResult := make(map[string]interface{})
	jsonResult["success"] = false

	var req model.HeartbeatRequest

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); nil != err {
		logger.Errorln("parseForm error", err)
		jsonResult["message"] = fmt.Sprintf("parseForm error: %v", err.Error())
	} else if "" == req.Hostname {
		jsonResult["message"] = "hostname is blank"
	} else {
		logger.Debugln("Heartbeat Request===>>>", req)
		strIp := this.Ctx.Request.Header.Get("X-Forwarded-For")
		if "" == strIp {
			strIp = this.Ctx.Request.Header.Get("Proxy-Client-IP")
		}
		if "" == strIp {
			strIp = this.Ctx.Request.Header.Get("WL-Proxy-Client-IP")
		}
		if "" == strIp {
			ip, _, err := net.SplitHostPort(this.Ctx.Request.RemoteAddr)
			if nil != err {
				logger.Errorln("SplitHostPort error", this.Ctx.Request.RemoteAddr)
			} else {
				strIp = ip
			}
		}

		go store.ParseHeartbeatRequest(&req, strIp)

		resp := model.HeartbeatResponse{
			ErrorMessage:  "",
			DesiredAgents: store.HostAgents.Status(req.Hostname),
		}
		logger.Debugln("Heartbeat Response<<<===", resp)

		this.Data["json"] = &resp
		jsonResult["success"] = true
	}

	if false == jsonResult["success"] {
		this.Data["json"] = &jsonResult
	}

	this.ServeJSON()
}
