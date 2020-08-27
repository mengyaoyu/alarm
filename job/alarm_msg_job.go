package job

import (
	"alarm/common"
	"alarm/models"
	"alarm/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
)

func InitAlarmMsgJob() {
	taskId := "SendAlarmMsg"
	taskFunc := func() {
		o := orm.NewOrm()
		dingTalkList := models.GetDingTalkList(o)

		for _, dingTalk := range dingTalkList {
			alarmMsgList := models.GetAlarmMsgListByAccessToken(dingTalk.AccessToken, 0, o)
			var ids = make([]int64, len(alarmMsgList))
			for idx, alarmMsg := range alarmMsgList {
				msg := alarmMsg.Msg + " 发生时间：" + alarmMsg.CreateTime.Format("2006-01-02 15:04:05")
				dingTalkMsg := `{"msgtype": "text", "text": {"content": "【水桥】` + msg + `"}}`
				logs.Info("发送钉钉机器人 %s 消息：%s", alarmMsg.AccessToken, dingTalkMsg)
				go utils.PostJson("https://oapi.dingtalk.com/robot/send?access_token="+alarmMsg.AccessToken, dingTalkMsg, "application/json")
				ids[idx] = alarmMsg.Id
			}
			models.BatchUpdateAlarmMsgSuccess(ids, o)
		}
	}

	if err := common.AlarmCronTab.AddByFunc(taskId, "0/1 * * * ?", taskFunc); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}

}
