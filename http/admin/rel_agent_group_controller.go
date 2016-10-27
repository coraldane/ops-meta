package admin

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
)

type RelAgentGroupController struct {
	beego.Controller
}

func (this *RelAgentGroupController) Get() {
	strAgentId := this.Ctx.Input.Param(":agentId")
	if "" == strAgentId {
		this.Data["Error"] = "未关联Agent"
	} else {
		agentId, _ := strconv.ParseInt(strAgentId, 10, 64)
		this.Data["agentId"] = agentId

		if 0 < agentId {
			list, err := models.QueryRelAgentGroupList(agentId)
			if nil != err {
				this.Data["Error"] = err.Error()
			} else {
				this.Data["items"] = list
			}
		}
	}
	this.TplName = "admin/rel_agent_group_main.html"
}

func (this *RelAgentGroupController) Save() {
	jsonResult := g.JsonResult{}

	item := models.RelAgentGroup{}
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

func (this *RelAgentGroupController) Delete() {
	jsonResult := g.JsonResult{}

	id, _ := this.GetInt64("Id")
	item := models.RelAgentGroup{Id: id}
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
