package models

import "github.com/astaxie/beego/orm"

type AlarmNoticeMonitorSql struct {
	Id        int64
	Cron      string
	Sql       string
	AliasName string
	Status    int
	DbId      int
	Note      string
	Sn        string
	ShowSql   int
}

func AddMonitorSql(cron string, sql string, dbId int) (int64, error) {

	m := new(AlarmNoticeMonitorSql)
	m.Status = 0
	m.Cron = cron
	m.Sql = sql
	o := orm.NewOrm()
	var db AlarmDbConnection
	_ = o.QueryTable("alarm_db_connection").Filter("id", dbId).One(&db)

	m.AliasName = db.AliasName
	m.DbId = db.Id

	return o.Insert(m)
}

func GetAlarmNoticeMonitorSqlList(status interface{}) []AlarmNoticeMonitorSql {
	var sqlList []AlarmNoticeMonitorSql
	o := orm.NewOrm()
	qs := o.QueryTable("alarm_notice_monitor_sql")
	if status != nil {
		qs.Filter("status", status).All(&sqlList)
	} else {
		qs.All(&sqlList)
	}
	return sqlList
}

func GetAlarmNoticeMonitorSqlById(id int64, o orm.Ormer) AlarmNoticeMonitorSql {
	var sql AlarmNoticeMonitorSql
	if o == nil {
		o = orm.NewOrm()
	}
	qs := o.QueryTable("alarm_notice_monitor_sql")
	qs.Filter("id", id).All(&sql)
	return sql
}
