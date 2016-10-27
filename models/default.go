package models

import (
	"github.com/astaxie/beego/orm"
)

func Init() {
	// register model
	orm.RegisterModelWithPrefix("t_", new(Endpoint), new(HostGroup), new(User),
		new(RelEndpointGroup), new(Agent), new(EndpointAgent), new(RelAgentGroup))
}
