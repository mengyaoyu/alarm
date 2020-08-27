package controllers

import (
	"alarm/dto"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AlarmMsgController struct {
	beego.Controller
}

// @router /alarm/msg/save [post]
func (c *AlarmMsgController) SaveAlarmMsg() {
	var msg dto.SaveAlarmMsg
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logs.Info(msg)
	}
	logs.Info(msg)
	result := map[string]interface{}{"code": "200"}

	c.Data["json"] = result
	c.ServeJSON()
}
