package services

import (
	"alarm/models"
	"alarm/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
)

func InitMonitorTask(cronTab *utils.Crontab) {

	sqlList := models.GetMonitorSqlList(1)

	for _, sql := range sqlList {
		taskId := "task" + strconv.FormatInt(sql.Id, 10)
		qsl := sql.Sql
		aliasName := sql.AliasName
		note := sql.Note
		showSql := sql.ShowSql
		accessToken := sql.AccessToken
		// 添加函数作为定时任务
		taskFunc := func() {
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

				dingTalkMsg := `{"msgtype": "text", "text": {"content": "【水桥】` + msg + `"}}`
				utils.PostJson("https://oapi.dingtalk.com/robot/send?access_token="+accessToken, dingTalkMsg, "application/json")
			}
		}
		logs.Info(taskId)
		if err := cronTab.AddByFunc(taskId, sql.Cron, taskFunc); err != nil {
			fmt.Printf("error to add crontab task:%s", err)
			os.Exit(-1)
		}

	}
}
