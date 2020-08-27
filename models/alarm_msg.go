package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	"time"
)

type AlarmMsg struct {
	Id          int64
	Msg         string
	Status      int
	AccessToken string
	CreateTime  time.Time
	UpdateTime  time.Time
}

func GetAlarmMsgList(status interface{}, o orm.Ormer) []AlarmMsg {
	var msgList []AlarmMsg
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

func GetAlarmMsgListByAccessToken(accessToken string, status interface{}, o orm.Ormer) []AlarmMsg {
	var msgList []AlarmMsg
	if o == nil {
		o = orm.NewOrm()
	}
	qs := o.QueryTable("alarm_msg")
	if o != nil {
		qs.Filter("status", status).Filter("access_token", accessToken).OrderBy("id").Limit(20).All(&msgList)
	} else {
		qs.Filter("access_token", accessToken).OrderBy("id").Limit(20).All(&msgList)

	}
	return msgList
}

func UpdateAlarmMsgSuccess(msg AlarmMsg, o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	_, err := o.Raw("update alarm_msg set status = 1 ,update_time = ? where id = ? and status = 0 ", time.Now(), msg.Id).Exec()
	if err != nil {
		log.Error("更新alarmMsg为通知成功失败 id: %s", msg.Id)
	}
}

func BatchUpdateAlarmMsgSuccess(ids []int64, o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	p, _ := o.Raw("update alarm_msg set status = 1 ,update_time = CURRENT_TIMESTAMP where id = ? and status = 0 ").Prepare()

	for _, id := range ids {
		p.Exec(id)
	}
	p.Close()

}
