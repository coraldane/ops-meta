package admin

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
	"github.com/coraldane/ops-meta/store"
)

type AgentController struct {
	beego.Controller
}

func (this *AgentController) Get() {
	this.Layout = "layout/default_layout.html"
	this.TplName = "admin/agent_main.html"
}

func (this *AgentController) List() {
	this.Data["json"] = store.DesiredAgentsMap
	this.ServeJSON()
}

func (this *AgentController) Query() {
	jsonResult := make(map[string]interface{})

	pageInfo := models.NewPageInfo(this.Ctx.Request)

	queryDto := models.QueryAgentDto{}
	if err := this.ParseForm(&queryDto); nil != err {
		logger.Errorln("parseForm error", err)
	}
	list, pageInfo := models.QueryAgentList(queryDto, pageInfo)

	jsonResult["items"] = list
	jsonResult["total"] = pageInfo.RowCount

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}

func (this *AgentController) Edit() {
	item := models.Agent{Name: this.GetString("agentName")}
	entity, _ := item.Get()

	this.Data["entity"] = entity
	this.TplName = "admin/agent_edit.html"
}

func (this *AgentController) Save() {
	jsonResult := g.JsonResult{}

	item := models.Agent{}
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

func (this *AgentController) Delete() {
	jsonResult := g.JsonResult{}

	id, _ := this.GetInt64("Id")
	item := models.Agent{Id: id, Name: this.GetString("name")}
	result, err := item.DeleteByPK()
	if nil != err {
		jsonResult.Message = err.Error()
	} else if result > 0 {
		store.DesiredAgentsMap.Delete(item.Name)
		jsonResult.Success = true
	} else {
		jsonResult.Message = "更新失败，数据不存在"
	}

	this.Data["json"] = &jsonResult
	this.ServeJSON()
}
