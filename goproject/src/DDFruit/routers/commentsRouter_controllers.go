package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"] = append(beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "ShowLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"] = append(beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "HandleLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"] = append(beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "ShowRegister",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"] = append(beego.GlobalControllerRouter["DDFruit/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "HandleRegister",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
