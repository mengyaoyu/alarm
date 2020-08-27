package models

import "github.com/astaxie/beego/orm"

type DbConnection struct {
	Id         int
	AliasName  string
	DriverName string
	DataSource string
}

func GetDbConnectionList() []DbConnection {

	var dbList []DbConnection

	o := orm.NewOrm()

	qs := o.QueryTable("db_connection")

	qs.All(&dbList)

	return dbList
}
