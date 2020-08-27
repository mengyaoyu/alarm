package models

import "github.com/astaxie/beego/orm"

type MsgTopic struct {
	Id int
	Sn string
}

func GetMsgTopicBySn(sn string, o orm.Ormer) MsgTopic {

	var msg MsgTopic

	if o == nil {
		o = orm.NewOrm()
	}
	_ = o.QueryTable("msg_topic").Filter("sn", sn).One(&msg)

	return msg
}
