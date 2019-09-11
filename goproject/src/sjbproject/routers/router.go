package routers

import (
	"github.com/astaxie/beego"
	"sjbproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/other", &controllers.IndexController{}, "get:GetMethod;post:PostMethod")

	beego.Router("/api/?:id", &controllers.ApiController{}, "get:GetApi")

	beego.Router("/num/?:id([0-9]+)", &controllers.NumController{}, "get:GetNum")

}
