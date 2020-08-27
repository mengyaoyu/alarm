package routers

import (
	"alarm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.AlarmMsgController{})
}
