package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
	"github.com/coraldane/ops-meta/utils"
)

type Endpoint struct {
	Id             int64
	GmtCreate      time.Time `orm:"type(datetime)"`
	GmtModified    time.Time `orm:"type(datetime)"`
	Hostname       string    `form:"hostname"`
	Ip             string    `form:"ip"`
	UpdaterVersion string    `form:"updaterVersion" orm:"null"`
	RunUser        string    `form:"runUser" orm:"null"`
}

func (this *Endpoint) TableUnique() [][]string {
	return [][]string{
		[]string{"Hostname"},
	}
}

func (this *Endpoint) Insert() (int64, error) {
	num, err := db.NewOrm().QueryTable(Endpoint{}).Filter("Hostname", this.Hostname).Count()
	if nil != err {
		return 0, err
	} else if 0 == num {
		this.GmtCreate = time.Now()
		this.GmtModified = time.Now()
		return db.NewOrm().Insert(this)
	} else {
		strSql := `update t_endpoint set gmt_modified=?, ip = ?, updater_version = ?, run_user = ? where hostname=?`
		result, err := db.NewOrm().Raw(strSql, utils.FormatUTCTime(time.Now()), this.Ip, this.UpdaterVersion, this.RunUser, this.Hostname).Exec()
		if nil != err {
			logger.Errorln("insert endpoint fail: ", err)
			return 0, err
		}
		return result.RowsAffected()
	}
}

func (this *Endpoint) DeleteByPK() (int64, error) {
	result, err := this.DeleteByCond()

	ea := EndpointAgent{Hostname: this.Hostname}
	ea.Delete([]string{})
	return result, err
}

func (this *Endpoint) DeleteByCond() (int64, error) {
	query := db.NewOrm().QueryTable(Endpoint{})
	if 0 < this.Id {
		query = query.Filter("Id", this.Id)
	}
	if "" != this.Hostname {
		query = query.Filter("Hostname", this.Hostname)
	}
	if "" != this.Ip {
		query = query.Filter("Ip", this.Ip)
	}
	return query.Delete()
}

func QueryEndpointList(queryDto QueryEndpointDto, pageInfo *PageInfo) ([]Endpoint, *PageInfo) {
	var rows []Endpoint
	query := db.NewOrm().QueryTable(Endpoint{})
	if "" != queryDto.Hostname {
		query = query.Filter("hostname__icontains", queryDto.Hostname)
	}
	if "" != queryDto.Ip {
		query = query.Filter("ip__contains", queryDto.Ip)
	}

	rowCount, err := query.Count()
	if nil != err {
		logger.Errorln("queryCount error", err)
		pageInfo.SetRowCount(0)
		return nil, pageInfo
	}
	pageInfo.SetRowCount(rowCount)

	_, err = query.OrderBy("-GmtModified").Offset(pageInfo.GetStartIndex()).Limit(pageInfo.PageSize).All(&rows)
	if nil != err {
		logger.Errorln("QueryEndpointList error", err)
	}
	return rows, pageInfo
}
