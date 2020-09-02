package routers

import (
	"alarm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	alarm := beego.NewNamespace("alarm-api",
		beego.NSInclude(&controllers.TestController{}),
		beego.NSInclude(&controllers.AlarmMsgController{}),
		beego.NSInclude(&controllers.MonitorSqlController{}),
	)
	beego.AddNamespace(alarm)
}
