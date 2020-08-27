package job

import (
	"alarm/common"
	"alarm/dto"
	"alarm/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"time"
)

func InitMonitorJob() {

	sqlList := models.GetMonitorSqlList(1)

	for _, sql := range sqlList {
		addMonitorSqlJob(sql)
	}
}

func addMonitorSqlJob(sql models.MonitorSql) {
	id := strconv.FormatInt(sql.Id, 10)
	taskId := common.JobPrefix + id
	qsl := sql.Sql
	aliasName := sql.AliasName
	note := sql.Note
	showSql := sql.ShowSql
	// 添加函数作为定时任务
	taskFunc := func() {
		logs.Info("run monitor sql : %s", qsl)
		var maps []orm.Params
		o := orm.NewOrm()
		o.Using(aliasName)
		_, _ = o.Raw(qsl).Values(&maps)

		result := maps[0]["cnt"].(string)
		int, _ := strconv.Atoi(result)
		if int > 0 {
			var msg string

			if showSql == 1 {
				msg = note + " : " + result + " 查询SQL: " + qsl
			} else {
				msg = note + " : " + result
			}
			requestNo := "monitor_sql_id_" + id + "_" + strconv.FormatInt(time.Now().Unix(), 10)
			saveAlarmMsg := dto.SaveAlarmMsg{Msg: msg, RequestNo: requestNo, Sn: sql.Sn}
			dto.DealSaveAlarmMsg(saveAlarmMsg)

		}
	}
	logs.Info(taskId)
	if err := common.AlarmCronTab.AddByFunc(taskId, sql.Cron, taskFunc); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}
}

func AddMonitorJob(sqlId int64) {
	o := orm.NewOrm()
	sql := models.GetMonitorSqlById(sqlId, o)

	addMonitorSqlJob(sql)
}
