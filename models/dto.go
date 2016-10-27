package models

import (
	"time"
)

type RelAgentGroupDto struct {
	Id          int64
	GmtCreate   time.Time
	GmtModified time.Time
	AgentId     int64
	HostGroupId int64
	GroupName   string
}
