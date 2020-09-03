package models

import "github.com/astaxie/beego/orm"

type AlarmDbConnection struct {
	Id         int
	AliasName  string
	DriverName string
	DataSource string
}

func GetAlarmDbConnectionList() []AlarmDbConnection {

	var dbList []AlarmDbConnection

	o := orm.NewOrm()

	_, _ = o.QueryTable("alarm_db_connection").All(&dbList)

	return dbList
}
