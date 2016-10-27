package admin

import (
	"github.com/astaxie/beego"
	// "github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/models"
	"github.com/coraldane/ops-meta/store"
)

type EndpointAgentController struct {
	beego.Controller
}

func (this *EndpointAgentController) Get() {
	hostname := this.Ctx.Input.Param(":hostname")
	if "" == hostname {
		this.Data["Error"] = "请选择Endpoint"
	} else {
		this.Data["DesiredAgentList"] = store.HostAgents.Status(hostname)
		this.Data["RealAgentList"] = models.QueryEndpointAgentList(hostname)
	}
	this.TplName = "admin/endpoint_agent.html"
}
