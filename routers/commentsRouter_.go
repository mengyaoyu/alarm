package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["alarm/controllers:AlarmMsgController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmMsgController"],
		beego.ControllerComments{
			Method:           "SaveAlarmMsg",
			Router:           "/alarm/msg/save",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:TestController"] = append(beego.GlobalControllerRouter["alarm/controllers:TestController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/test/get",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
