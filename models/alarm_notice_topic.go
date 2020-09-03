package models

import "github.com/astaxie/beego/orm"

type AlarmNoticeTopic struct {
	Id int
	Sn string
}

func GetAlarmNoticeTopicBySn(sn string, o orm.Ormer) AlarmNoticeTopic {

	var msg AlarmNoticeTopic

	if o == nil {
		o = orm.NewOrm()
	}
	_ = o.QueryTable("alarm_notice_topic").Filter("sn", sn).One(&msg)

	return msg
}
