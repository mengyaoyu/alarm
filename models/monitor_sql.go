package models

import "github.com/astaxie/beego/orm"

type MonitorSql struct {
	Id          int64
	Cron        string
	Sql         string
	AliasName   string
	Status      int
	DbId        int
	Note        string
	AccessToken string
	ShowSql     int
}

func AddMonitorSql(cron string, sql string, dbId int) (int64, error) {

	m := new(MonitorSql)
	m.Status = 0
	m.Cron = cron
	m.Sql = sql
	o := orm.NewOrm()
	var db DbConnection
	o.QueryTable("db_connection").Filter("id", dbId).One(&db)

	m.AliasName = db.AliasName
	m.DbId = db.Id

	return o.Insert(m)
}

func GetMonitorSqlList(status interface{}) []MonitorSql {
	var sqlList []MonitorSql
	o := orm.NewOrm()
	qs := o.QueryTable("monitor_sql")
	if status != nil {
		qs.Filter("status", status)
	}
	qs.All(&sqlList)
	return sqlList
}
