package admin

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
)

type HostGroupController struct {
	beego.Controller
}

func (this *HostGroupController) Get() {

	this.Layout = "layout/default_layout.html"
	this.TplName = "admin/host_group_main.html"
}

func (this *HostGroupController) Query() {
	jsonResult := make(map[string]interface{})

	pageInfo := models.NewPageInfo(this.Ctx.Request)

	queryDto := models.QueryHostGroupDto{}
	if err := this.ParseForm(&queryDto); nil != err {
		logger.Errorln("parseForm error", err)
	}
	list, pageInfo := models.QueryHostGroupList(queryDto, pageInfo)

	jsonResult["items"] = list
	jsonResult["total"] = pageInfo.RowCount

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}

func (this *HostGroupController) Edit() {
	item := models.HostGroup{GroupName: this.GetString("name")}
	entity, _ := item.Get()

	this.Data["entity"] = entity
	this.TplName = "admin/host_group_edit.html"
}

func (this *HostGroupController) Save() {
	jsonResult := g.JsonResult{}

	item := models.HostGroup{}
	if err := this.ParseForm(&item); nil != err {
		logger.Errorln("parseForm error", err)
		jsonResult.Message = err.Error()
	} else {
		if 0 >= item.Id {
			item.GmtCreate = time.Now()
		}
		result, err := item.Insert()
		if nil != err {
			jsonResult.Message = err.Error()
		} else if 0 < result {
			jsonResult.Success = true
		} else {
			jsonResult.Message = "保存失败，未知错误"
		}
	}

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}

func (this *HostGroupController) Delete() {
	jsonResult := g.JsonResult{}

	id, _ := this.GetInt64("Id")
	item := models.HostGroup{Id: id}
	result, err := item.DeleteByPK()
	if nil != err {
		jsonResult.Message = err.Error()
	} else if result > 0 {
		jsonResult.Success = true
	} else {
		jsonResult.Message = "更新失败，数据不存在"
	}

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}
