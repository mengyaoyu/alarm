package controllers

import (
	"alarm/common"
	"alarm/dto"
	"alarm/job"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type MonitorSqlController struct {
	beego.Controller
}

// @router /monitor/start/task [post]
func (c *MonitorSqlController) StartMonitorTask() {
	var task dto.MonitorJobOnOff
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &task)
	if err != nil {
		logs.Info(task)
	}

	cacheKey := "StartMonitorTask_" + strconv.FormatInt(task.Id, 10)

	result := map[string]interface{}{"code": "200"}

	if !common.CacheMemory.IsExist(cacheKey) {
		common.CacheMemory.Put(cacheKey, cacheKey, 30*time.Second)
		job.AddMonitorJob(task.Id)
		logs.Info("StartMonitorTask %s", task.Id)
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /monitor/stop/task [post]
func (c *MonitorSqlController) StopMonitorTask() {
	var task dto.MonitorJobOnOff
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &task)
	if err != nil {
		logs.Info(task)
	}
	common.AlarmCronTab.DelByID(common.JobPrefix + strconv.FormatInt(task.Id, 10))

	result := map[string]interface{}{"code": "200"}

	c.Data["json"] = result
	c.ServeJSON()
}
