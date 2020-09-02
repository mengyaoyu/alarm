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
	c.Data["json"] = map[string]interface{}{"code": "200"}
	c.ServeJSON()
}
