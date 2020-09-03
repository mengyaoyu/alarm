package controllers

import (
	"alarm/common"
	"alarm/dto"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type AlarmNoticeMsgController struct {
	beego.Controller
}

// @router /alarm/msg/save [post]
func (c *AlarmNoticeMsgController) SaveAlarmMsg() {
	var msg dto.SaveAlarmNoticeMsg
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logs.Info(msg)
	}

	cacheKey := "SaveAlarmMsg_" + msg.RequestNo

	result := map[string]interface{}{"code": "200"}

	if !common.CacheMemory.IsExist(cacheKey) {
		common.CacheMemory.Put(cacheKey, cacheKey, 30*time.Second)
		dto.DealSaveAlarmNoticeMsg(msg)
	}

	logs.Info(msg)

	c.Data["json"] = result
	c.ServeJSON()
}
