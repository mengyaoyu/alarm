package job

import (
	"alarm/common"
	"alarm/models"
	"alarm/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
)

func InitAlarmNoticeMsgJob() {
	taskId := "SendAlarmMsg"
	taskFunc := func() {
		o := orm.NewOrm()
		dingTalkList := models.GetAlarmNoticeListenerList(o, "DT")

		for _, dingTalk := range dingTalkList {
			alarmMsgList := models.GetAlarmNoticeListenerListByAccessToken(dingTalk.AccessToken, 0, o)
			var ids = make([]int64, len(alarmMsgList))
			for idx, alarmMsg := range alarmMsgList {
				msg := alarmMsg.Msg + " 时间:" + alarmMsg.CreateTime.In(common.CstSh).Format("2006-01-02 15:04:05")

				text := map[string]string{
					"content": msg,
				}

				msgMap := map[string]interface{}{
					"msgtype": "text",
					"text":    text,
				}

				msgJson, _ := json.Marshal(msgMap)
				dingTalkMsg := string(msgJson)

				logs.Info("发送钉钉机器人 %s 消息：%s", alarmMsg.AccessToken, dingTalkMsg)
				go utils.PostJson("https://oapi.dingtalk.com/robot/send?access_token="+alarmMsg.AccessToken, dingTalkMsg, "application/json")
				ids[idx] = alarmMsg.Id
			}
			models.BatchUpdateAlarmNoticeMsgSuccess(ids, o)
		}
	}

	if err := common.AlarmCronTab.AddByFunc(taskId, "0/1 * * * ?", taskFunc); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}

}
