package models

import (
	"github.com/astaxie/beego/orm"
)

type DingTalk struct {
	Id          int
	AccessToken string
}

func GetAccessTokenByMsgTopicSn(sn string, o orm.Ormer) []DingTalk {
	if o == nil {
		o = orm.NewOrm()
	}
	var dingTalk []DingTalk

	_, _ = o.Raw("SELECT dt.id,dt.access_token FROM ding_talk dt INNER JOIN msg_topic_ding_talk tdt ON dt.id = tdt.ding_talk_id INNER JOIN msg_topic t on tdt.topic_id = t.id WHERE sn = ?", sn).QueryRows(&dingTalk)

	return dingTalk
}

func GetDingTalkList(o orm.Ormer) []DingTalk {
	if o == nil {
		o = orm.NewOrm()
	}
	var dingTalk []DingTalk

	o.QueryTable("ding_talk").All(&dingTalk)

	return dingTalk
}
