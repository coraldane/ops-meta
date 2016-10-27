package admin

import (
	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
	"github.com/coraldane/ops-meta/store"
)

type EndpointController struct {
	beego.Controller
}

func (this *EndpointController) Get() {

	this.Layout = "layout/default_layout.html"
	this.TplName = "admin/endpoint_main.html"
}

func (this *EndpointController) Query() {
	jsonResult := make(map[string]interface{})

	pageInfo := models.NewPageInfo(this.Ctx.Request)

	queryDto := models.QueryEndpointDto{}
	if err := this.ParseForm(&queryDto); nil != err {
		logger.Errorln("parseForm error", err)
	}
	list, pageInfo := models.QueryEndpointList(queryDto, pageInfo)

	jsonResult["items"] = list
	jsonResult["total"] = pageInfo.RowCount

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}

func (this *EndpointController) Delete() {
	jsonResult := g.JsonResult{}

	id, _ := this.GetInt64("Id")
	hostname := this.GetString("Hostname")
	item := models.Endpoint{Id: id}
	result, err := item.DeleteByPK()
	if nil != err {
		jsonResult.Message = err.Error()
	} else if result > 0 {
		store.HostAgents.Delete(hostname)
		jsonResult.Success = true
	} else {
		jsonResult.Message = "更新失败，数据不存在"
	}

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}
