package models

import (
	"alarm/common"
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	"time"
)

type AlarmNoticeMsg struct {
	Id          int64
	Msg         string
	Status      int
	AccessToken string
	CreateTime  time.Time
	UpdateTime  time.Time
	NoticeType  string
}

func GetAlarmNoticeMsgList(status interface{}, o orm.Ormer) []AlarmNoticeMsg {
	var msgList []AlarmNoticeMsg
	if o == nil {
		o = orm.NewOrm()
	}
	qs := o.QueryTable("alarm_msg")
	if status != nil {
		qs.Filter("status", status).All(&msgList)
	} else {
		qs.All(&msgList)
	}
	return msgList
}

func GetAlarmNoticeListenerListByAccessToken(accessToken string, status interface{}, o orm.Ormer) []AlarmNoticeMsg {
	var msgList []AlarmNoticeMsg
	if o == nil {
		o = orm.NewOrm()
	}
	qs := o.QueryTable("alarm_notice_msg")
	if o != nil {
		qs.Filter("status", status).Filter("access_token", accessToken).OrderBy("id").Limit(20).All(&msgList)
	} else {
		qs.Filter("access_token", accessToken).OrderBy("id").Limit(20).All(&msgList)

	}
	return msgList
}

func UpdateAlarmNoticeMsgSuccess(msg AlarmNoticeMsg, o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	_, err := o.Raw("update alarm_notice_msg set status = 1 ,update_time = ? where id = ? and status = 0 ", time.Now().In(common.CstSh), msg.Id).Exec()
	if err != nil {
		log.Error("更新alarmMsg为通知成功失败 id: %s", msg.Id)
	}
}

func BatchUpdateAlarmNoticeMsgSuccess(ids []int64, o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	p, _ := o.Raw("update alarm_notice_msg set status = 1 ,update_time = CURRENT_TIMESTAMP where id = ? and status = 0 ").Prepare()

	for _, id := range ids {
		p.Exec(id)
	}
	p.Close()

}
