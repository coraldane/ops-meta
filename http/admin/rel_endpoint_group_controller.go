package admin

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
)

type RelEndpointGroupController struct {
	beego.Controller
}

func (this *RelEndpointGroupController) Get() {
	strHostGroupId := this.Ctx.Input.Param(":hostGroupId")
	if "" == strHostGroupId {
		this.Data["Error"] = "未关联HostGroup"
	} else {
		hostGroupId, _ := strconv.Atoi(strHostGroupId)
		this.Data["hostGroupId"] = hostGroupId
	}

	this.TplName = "admin/rel_endpoint_group_main.html"
}

func (this *RelEndpointGroupController) Query() {
	jsonResult := make(map[string]interface{})

	pageInfo := models.NewPageInfo(this.Ctx.Request)

	queryDto := models.QueryRelEndpointGroupDto{}
	if err := this.ParseForm(&queryDto); nil != err {
		logger.Errorln("parseForm error", err)
	}
	list, pageInfo := models.QueryRelEndpointGroupList(queryDto, pageInfo)

	jsonResult["items"] = list
	jsonResult["total"] = pageInfo.RowCount

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}

func (this *RelEndpointGroupController) Edit() {
	id, _ := this.GetInt64("Id")
	item := models.RelEndpointGroup{Id: id}
	entity, _ := item.Get()

	this.Data["entity"] = entity
	this.TplName = "admin/rel_endpoint_group_edit.html"
}

func (this *RelEndpointGroupController) Save() {
	jsonResult := g.JsonResult{}

	item := models.RelEndpointGroup{}
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

func (this *RelEndpointGroupController) Delete() {
	jsonResult := g.JsonResult{}

	id, _ := this.GetInt64("Id")
	item := models.RelEndpointGroup{Id: id}
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
