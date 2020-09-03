package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
        beego.ControllerComments{
            Method: "StartMonitorTask",
            Router: "/monitor/start/task",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
        beego.ControllerComments{
            Method: "StopMonitorTask",
            Router: "/monitor/stop/task",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMsgController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMsgController"],
        beego.ControllerComments{
            Method: "SaveAlarmMsg",
            Router: "/alarm/msg/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["alarm/controllers:TestController"] = append(beego.GlobalControllerRouter["alarm/controllers:TestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/test/get",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
