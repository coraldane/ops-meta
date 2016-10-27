package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
)

/**
RelType - fixed 固定值, regex 正则
*/
type RelEndpointGroup struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModified time.Time `orm:"auto_now_add;type(datetime)"`
	HostGroupId int64     `form:"hostGroupId"`
	RelType     string    `form:"relType"`
	PropName    string    `form:"propName"`
	PropValue   string    `form:"propValue"`
}

func (this *RelEndpointGroup) Insert() (int64, error) {
	this.GmtModified = time.Now()

	if 0 < this.Id {
		return db.NewOrm().Update(this)
	} else {
		this.GmtCreate = time.Now()
		return db.NewOrm().Insert(this)
	}
}

func (this *RelEndpointGroup) Get() (*RelEndpointGroup, error) {
	var retValue RelEndpointGroup
	err := db.NewOrm().QueryTable(RelEndpointGroup{}).Filter("Id", this.Id).One(&retValue)
	return &retValue, err
}

func (this *RelEndpointGroup) DeleteByPK() (int64, error) {
	result, err := this.DeleteByCond()
	return result, err
}

func (this *RelEndpointGroup) DeleteByCond() (int64, error) {
	query := db.NewOrm().QueryTable(RelEndpointGroup{})
	if 0 < this.Id {
		query = query.Filter("Id", this.Id)
	}
	if 0 < this.HostGroupId {
		query = query.Filter("HostGroupId", this.HostGroupId)
	}
	if "" != this.RelType {
		query = query.Filter("RelType", this.RelType)
	}
	if "" != this.PropName {
		query = query.Filter("PropName", this.PropName)
	}
	if "" != this.PropValue {
		query = query.Filter("PropValue", this.PropValue)
	}
	return query.Delete()
}

func QueryRelEndpointGroupList(queryDto QueryRelEndpointGroupDto, pageInfo *PageInfo) ([]RelEndpointGroup, *PageInfo) {
	var rows []RelEndpointGroup
	query := db.NewOrm().QueryTable(RelEndpointGroup{})
	if 0 < queryDto.HostGroupId {
		query = query.Filter("HostGroupId", queryDto.HostGroupId)
	}
	if "" != queryDto.RelType {
		query = query.Filter("RelType", queryDto.RelType)
	}
	if "" != queryDto.PropName {
		query = query.Filter("PropName", queryDto.PropName)
	}
	if "" != queryDto.PropValue {
		query = query.Filter("PropValue", queryDto.PropValue)
	}

	rowCount, err := query.Count()
	if nil != err {
		logger.Errorln("queryCount error", err)
		pageInfo.SetRowCount(0)
		return nil, pageInfo
	}
	pageInfo.SetRowCount(rowCount)

	_, err = query.OrderBy("RelType").Offset(pageInfo.GetStartIndex()).Limit(pageInfo.PageSize).All(&rows)
	if nil != err {
		logger.Errorln("QueryRelEndpointGroupList error", err)
	}
	return rows, pageInfo
}
