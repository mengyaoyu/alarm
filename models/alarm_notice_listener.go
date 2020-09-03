package models

import (
	"github.com/astaxie/beego/orm"
)

type AlarmNoticeListener struct {
	Id          int
	AccessToken string
	NoticeType  string
}

func GetAccessTokenByAlarmNoticeTopicSn(sn string, o orm.Ormer) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}
	var noticeListeners []AlarmNoticeListener

	_, _ = o.Raw("SELECT anl.id,anl.access_token,anl.notice_type FROM alarm_notice_listener anl "+
		"INNER JOIN alarm_notice_topic_listener antl ON anl.id = antl.listener_id  "+
		"INNER JOIN alarm_notice_topic ant ON ant.id = antl.topic_id where ant.sn = ?", sn).QueryRows(&noticeListeners)

	return noticeListeners
}

func GetAlarmNoticeListenerList(o orm.Ormer, noticeType string) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}
	var noticeListeners []AlarmNoticeListener

	o.QueryTable("alarm_notice_listener").Filter("notice_type", noticeType).All(&noticeListeners)

	return noticeListeners
}
