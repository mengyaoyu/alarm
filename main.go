package main

import (
	"alarm/models"
	_ "alarm/routers"
	"alarm/services"
	"alarm/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	initConfig()

	//logs.SetLogger(logs.AdapterConsole, `{"level":1}`)

	connection := beego.AppConfig.String("db.username") + ":" + beego.AppConfig.String("db.pwd") + "@tcp(" + beego.AppConfig.String("db.host") + ")/" + beego.AppConfig.String("db.name") + "?charset=utf8"

	logs.SetLogger(logs.AdapterFile, `{"filename":"info.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", connection)

	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 10)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(models.MonitorSql), new(models.DbConnection))

	dbList := models.GetDbConnectionList()

	for _, db := range dbList {
		orm.RegisterDataBase(db.AliasName, db.DriverName, db.DataSource)
		orm.SetMaxIdleConns(db.AliasName, 10)
		orm.SetMaxOpenConns(db.AliasName, 10)
	}

	cronTab := utils.NewCrontab()

	services.InitMonitorTask(cronTab)

	cronTab.Start()

	defer cronTab.Stop()

	logs.Info("start success !")
	beego.Run()
}

func initConfig() {

	rootPath := GetAPPRootPath()

	beego.SetViewsPath(rootPath + "/views")

	beego.LoadAppConfig("ini", rootPath+"/conf/app.conf")

	beego.SetStaticPath("static", rootPath+"/static")

}

func GetAPPRootPath() string {

	file, err := exec.LookPath(os.Args[0])

	if err != nil {

		return ""

	}

	p, err := filepath.Abs(file)

	if err != nil {

		return ""

	}

	return filepath.Dir(p)

}
