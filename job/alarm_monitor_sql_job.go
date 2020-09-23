package job

import (
	"alarm/common"
	"alarm/dto"
	"alarm/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"time"
)

func InitAlarmNoticeMonitorJob() {

	sqlList := models.GetAlarmNoticeMonitorSqlList(1)

	for _, sql := range sqlList {
		addAlarmNoticeMonitorSqlJob(sql)
	}
}

func addAlarmNoticeMonitorSqlJob(sql models.AlarmNoticeMonitorSql) {
	id := strconv.FormatInt(sql.Id, 10)
	taskId := common.JobPrefix + id
	qsl := sql.Sql
	aliasName := sql.AliasName
	note := sql.Note
	showSql := sql.ShowSql
	// 添加函数作为定时任务
	taskFunc := func() {
		var maps []orm.Params
		o := orm.NewOrm()
		o.Using(aliasName)
		_, _ = o.Raw(qsl).Values(&maps)

		logs.Info("run monitor sql : %s  result : %s", qsl, maps)
		if len(maps) > 0 {
			resultJson, _ := json.Marshal(maps)
			result := string(resultJson)

			var msg string
			if showSql == 1 {
				msg = note + result + " 查询SQL: " + qsl
			} else {
				msg = note + result
			}
			requestNo := "monitor_sql_id_" + taskId + "_" + strconv.FormatInt(time.Now().In(common.CstSh).Unix(), 10)
			saveAlarmMsg := dto.SaveAlarmNoticeMsg{Msg: msg, RequestNo: requestNo, Sn: sql.Sn}
			dto.DealSaveAlarmNoticeMsg(saveAlarmMsg)
		}
	}
	logs.Info(taskId)
	if err := common.AlarmCronTab.AddByFunc(taskId, sql.Cron, taskFunc); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}
}

func AddAlarmMonitorJob(sqlId int64) {
	o := orm.NewOrm()
	sql := models.GetAlarmNoticeMonitorSqlById(sqlId, o)

	addAlarmNoticeMonitorSqlJob(sql)
}
