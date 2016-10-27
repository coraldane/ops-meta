package models

import (
	"time"

	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
)

type HostGroup struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModified time.Time `orm:"auto_now_add;type(datetime)"`
	GroupName   string    `form:"groupName"`
}

func (this *HostGroup) TableUnique() [][]string {
	return [][]string{
		[]string{"GroupName"},
	}
}

func (this *HostGroup) Insert() (int64, error) {
	this.GmtModified = time.Now()

	if 0 < this.Id {
		return db.NewOrm().Update(this)
	} else {
		this.GmtCreate = time.Now()
		return db.NewOrm().Insert(this)
	}
}

func (this *HostGroup) Get() (*HostGroup, error) {
	var retValue HostGroup
	err := db.NewOrm().QueryTable(HostGroup{}).Filter("GroupName", this.GroupName).One(&retValue)
	return &retValue, err
}

func (this *HostGroup) CheckExists() bool {
	return db.NewOrm().QueryTable(HostGroup{}).Filter("GroupName", this.GroupName).Exist()
}

func (this *HostGroup) DeleteByPK() (int64, error) {
	result, err := this.DeleteByCond()
	return result, err
}

func (this *HostGroup) DeleteByCond() (int64, error) {
	query := db.NewOrm().QueryTable(HostGroup{})
	if 0 < this.Id {
		query = query.Filter("Id", this.Id)
	}
	if "" != this.GroupName {
		query = query.Filter("GroupName", this.GroupName)
	}
	return query.Delete()
}

func QueryHostGroupList(queryDto QueryHostGroupDto, pageInfo *PageInfo) ([]HostGroup, *PageInfo) {
	var rows []HostGroup
	query := db.NewOrm().QueryTable(HostGroup{})
	if "" != queryDto.GroupName {
		query = query.Filter("GroupName__contains", queryDto.GroupName)
	}

	rowCount, err := query.Count()
	if nil != err {
		logger.Errorln("queryCount error", err)
		pageInfo.SetRowCount(0)
		return nil, pageInfo
	}
	pageInfo.SetRowCount(rowCount)

	_, err = query.OrderBy("GroupName").Offset(pageInfo.GetStartIndex()).Limit(pageInfo.PageSize).All(&rows)
	if nil != err {
		logger.Errorln("QueryHostGroupList error", err)
	}
	return rows, pageInfo
}
