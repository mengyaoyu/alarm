package routers

import (
	"alarm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	alarm := beego.NewNamespace("alarm-api",
		beego.NSInclude(&controllers.TestController{}),
		beego.NSInclude(&controllers.AlarmNoticeMsgController{}),
		beego.NSInclude(&controllers.AlarmNoticeMonitorSqlController{}),
	)
	beego.AddNamespace(alarm)
}
