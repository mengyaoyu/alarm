package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// @router /test/get [get]
func (c *TestController) Get() {
	c.Ctx.WriteString(beego.AppConfig.String("test"))
	c.ServeJSONP()
}
