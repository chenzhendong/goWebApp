package controllers
import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.FinishRouter, globalOutBoundRule, true)
}

func globalOutBoundRule(ctx *context.Context)  {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
}
