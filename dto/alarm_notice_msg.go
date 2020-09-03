package dto

import (
	"alarm/common"
	"alarm/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type SaveAlarmNoticeMsg struct {
	Msg       string
	Sn        string
	RequestNo string
}

func DealSaveAlarmNoticeMsg(msg SaveAlarmNoticeMsg) {
	o := orm.NewOrm()

	alarmNoticeListenerList := models.GetAccessTokenByAlarmNoticeTopicSn(msg.Sn, o)

	var alarmMsgList = make([]models.AlarmNoticeMsg, len(alarmNoticeListenerList))

	for idx, listener := range alarmNoticeListenerList {
		alarmMsg := models.AlarmNoticeMsg{Msg: msg.Msg, Status: 0, AccessToken: listener.AccessToken,
			CreateTime: time.Now().In(common.CstSh), UpdateTime: time.Now().In(common.CstSh), NoticeType: listener.NoticeType}
		alarmMsgList[idx] = alarmMsg
	}

	o.InsertMulti(1, alarmMsgList)
}
