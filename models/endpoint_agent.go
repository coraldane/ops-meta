package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
)

type EndpointAgent struct {
	Id           int64
	GmtCreate    time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModified  time.Time `orm:"auto_now_add;type(datetime)"`
	Hostname     string    `json:"hostname"`
	AgentName    string    `json:"agentName"`
	AgentVersion string    `json:"agentVersion" orm:"null"`
	RunUser      string    `json:"runUser" orm:"null"`
	WorkDir      string    `json:"workDir" orm:"null"`
	Status       string    `json:"status" orm:"null"`
	GmtReport    time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *EndpointAgent) TableUnique() [][]string {
	return [][]string{
		[]string{"Hostname", "AgentName"},
	}
}

func (this *EndpointAgent) Insert() (int64, error) {
	this.GmtModified = time.Now()
	return db.NewOrm().InsertOrUpdate(this)
}

func (this *EndpointAgent) Delete(excludes []string) (int64, error) {
	query := db.NewOrm().QueryTable(EndpointAgent{}).Filter("HostName", this.Hostname)
	if 0 < len(excludes) {
		query = query.Exclude("agent_name__in", excludes)
	}
	return query.Delete()
}

func QueryEndpointAgentList(hostname string) []EndpointAgent {
	var rows []EndpointAgent
	_, err := db.NewOrm().QueryTable(EndpointAgent{}).Filter("Hostname", hostname).All(&rows)
	if nil != err {
		logger.Errorln("QueryEndpointAgentList error", err)
	}
	return rows
}
