package dto

import (
	"alarm/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type SaveAlarmMsg struct {
	Msg       string
	Sn        string
	RequestNo string
}

func DealSaveAlarmMsg(msg SaveAlarmMsg) {
	o := orm.NewOrm()

	dingTalkList := models.GetAccessTokenByMsgTopicSn(msg.Sn, o)

	var alarmMsgList = make([]models.AlarmMsg, len(dingTalkList))

	for idx, dingTalk := range dingTalkList {
		alarmMsg := models.AlarmMsg{Msg: msg.Msg, Status: 0, AccessToken: dingTalk.AccessToken, CreateTime: time.Now(), UpdateTime: time.Now()}
		alarmMsgList[idx] = alarmMsg
	}

	o.InsertMulti(1, alarmMsgList)
}
