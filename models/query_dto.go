package models

type QueryAgentDto struct {
	AgentName string `form:"agentName"`
	RunUser   string `form:"runUser"`
}

type QueryHostGroupDto struct {
	GroupName string `form:"groupName"`
}

type QueryRelEndpointGroupDto struct {
	HostGroupId int64  `form:"hostGroupId"`
	RelType     string `form:"relType"`
	PropName    string `form:"propName"`
	PropValue   string `form:"propValue"`
}

type QueryEndpointDto struct {
	Hostname string `form:"hostname"`
	Ip       string `form:"ip"`
}

type QueryUserDto struct {
	UserName string `form:"userName"`
	RealName string `form:"realName"`
	RoleName string `form:"roleName"`
}
