package admin

import (
	"github.com/astaxie/beego"
)

func ConfigRoutes() {
	beego.Router("/admin/agent", &AgentController{})
	beego.Router("/admin/agent/list", &AgentController{}, "get:List")
	beego.Router("/admin/agent/query", &AgentController{}, "post:Query")
	beego.Router("/admin/agent/edit", &AgentController{}, "get:Edit")
	beego.Router("/admin/agent/save", &AgentController{}, "post:Save")
	beego.Router("/admin/agent/delete", &AgentController{}, "post:Delete")

	beego.Router("/admin/endpoint", &EndpointController{})
	beego.Router("/admin/endpoint/query", &EndpointController{}, "post:Query")
	beego.Router("/admin/endpoint/delete", &EndpointController{}, "post:Delete")

	beego.Router("/admin/hostGroup", &HostGroupController{})
	beego.Router("/admin/hostGroup/query", &HostGroupController{}, "post:Query")
	beego.Router("/admin/hostGroup/edit", &HostGroupController{}, "get:Edit")
	beego.Router("/admin/hostGroup/save", &HostGroupController{}, "post:Save")
	beego.Router("/admin/hostGroup/delete", &HostGroupController{}, "post:Delete")

	beego.Router("/admin/relEndpointGroup/:hostGroupId", &RelEndpointGroupController{})
	beego.Router("/admin/relEndpointGroup/query", &RelEndpointGroupController{}, "post:Query")
	beego.Router("/admin/relEndpointGroup/edit", &RelEndpointGroupController{}, "get:Edit")
	beego.Router("/admin/relEndpointGroup/save", &RelEndpointGroupController{}, "post:Save")
	beego.Router("/admin/relEndpointGroup/delete", &RelEndpointGroupController{}, "post:Delete")

	beego.Router("/admin/relAgentGroup/:agentId", &RelAgentGroupController{})
	beego.Router("/admin/relAgentGroup/save", &RelAgentGroupController{}, "post:Save")
	beego.Router("/admin/relAgentGroup/delete", &RelAgentGroupController{}, "post:Delete")

	beego.Router("/admin/endpointAgent/:hostname", &EndpointAgentController{})
}
