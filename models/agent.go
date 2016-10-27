package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
)

type Agent struct {
	Id              int64
	GmtCreate       time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModified     time.Time `orm:"auto_now_add;type(datetime)"`
	Name            string
	Version         string
	Tarball         string `orm:"null"`
	Md5             string `orm:"null"`
	Cmd             string `orm:"null"`
	RunUser         string `orm:"null"`
	WorkDir         string `orm:"null"`
	ConfigFileName  string `orm:"null"`
	ConfigRemoteUrl string `orm:"null"`
}

func (this *Agent) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (this *Agent) Insert() (int64, error) {
	this.GmtModified = time.Now()
	return db.NewOrm().InsertOrUpdate(this)
}

func (this *Agent) Get() (*Agent, error) {
	var retValue Agent
	err := db.NewOrm().QueryTable(Agent{}).Filter("Name", this.Name).One(&retValue)
	return &retValue, err
}

func (this *Agent) CheckExists() bool {
	return db.NewOrm().QueryTable(Agent{}).Filter("Name", this.Name).Exist()
}

func (this *Agent) DeleteByPK() (int64, error) {
	result, err := this.DeleteByCond()

	relAgentGroup := RelAgentGroup{AgentId: this.Id}
	relAgentGroup.DeleteByCond()

	return result, err
}

func (this *Agent) DeleteByCond() (int64, error) {
	query := db.NewOrm().QueryTable(Agent{})
	if 0 < this.Id {
		query = query.Filter("Id", this.Id)
	}
	if "" != this.Name {
		query = query.Filter("Name", this.Name)
	}
	if "" != this.Version {
		query = query.Filter("Version", this.Version)
	}
	return query.Delete()
}

func QueryAgentList(queryDto QueryAgentDto, pageInfo *PageInfo) ([]Agent, *PageInfo) {
	var rows []Agent
	query := db.NewOrm().QueryTable(Agent{})
	if "" != queryDto.AgentName {
		query = query.Filter("Name", queryDto.AgentName)
	}
	if "" != queryDto.RunUser {
		query = query.Filter("RunUser", queryDto.RunUser)
	}

	rowCount, err := query.Count()
	if nil != err {
		logger.Errorln("queryCount error", err)
		pageInfo.SetRowCount(0)
		return nil, pageInfo
	}
	pageInfo.SetRowCount(rowCount)

	_, err = query.OrderBy("Name").Offset(pageInfo.GetStartIndex()).Limit(pageInfo.PageSize).All(&rows)
	if nil != err {
		logger.Errorln("QueryAgentList error", err)
	}
	return rows, pageInfo
}
