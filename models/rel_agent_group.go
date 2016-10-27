package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
)

type RelAgentGroup struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModified time.Time `orm:"auto_now_add;type(datetime)"`
	AgentId     int64     `form:"agentId"`
	HostGroupId int64     `form:"hostGroupId"`
}

func (this *RelAgentGroup) TableUnique() [][]string {
	return [][]string{
		[]string{"AgentId", "HostGroupId"},
	}
}

func (this *RelAgentGroup) Insert() (int64, error) {
	this.GmtModified = time.Now()

	if 0 < this.Id {
		return db.NewOrm().Update(this)
	} else {
		this.GmtCreate = time.Now()
		return db.NewOrm().Insert(this)
	}
}

func (this *RelAgentGroup) DeleteByPK() (int64, error) {
	result, err := this.DeleteByCond()
	return result, err
}

func (this *RelAgentGroup) DeleteByCond() (int64, error) {
	query := db.NewOrm().QueryTable(RelAgentGroup{})
	if 0 < this.Id {
		query = query.Filter("Id", this.Id)
	}
	if 0 < this.HostGroupId {
		query = query.Filter("HostGroupId", this.HostGroupId)
	}
	if 0 < this.AgentId {
		query = query.Filter("AgentId", this.AgentId)
	}
	return query.Delete()
}

func QueryRelAgentGroupList(agentId int64) ([]RelAgentGroupDto, error) {
	var rows []RelAgentGroupDto
	_, err := db.NewOrm().Raw("select t.id, t.gmt_create,t.gmt_modified, t.host_group_id, a.group_name from t_rel_agent_group t, t_host_group a where t.agent_id=? and t.host_group_id=a.id", agentId).QueryRows(&rows)
	if nil != err {
		logger.Errorln("QueryRelAgentGroupList error", err)
	}
	return rows, err
}
