package api

import (
	"github.com/astaxie/beego"
	"github.com/toolkits/file"

	"github.com/coraldane/ops-meta/g"
)

type FrameController struct {
	beego.Controller
}

func (this *FrameController) Version() {
	this.Ctx.WriteString(g.VERSION)
}

func (this *FrameController) WorkDir() {
	this.Ctx.WriteString(file.SelfDir())
}
