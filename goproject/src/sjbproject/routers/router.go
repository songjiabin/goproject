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

	beego.Router("/user/?:id([\\w]+)", &controllers.UserNameController{}, "get:GetUserName")

	beego.Router("/download/*.*", &controllers.DownLoadController{}, "get:DownLoadFile")

	beego.Router("/all/*", &controllers.AllController{}, "get:All")

	beego.Router("/mysql/", &controllers.MySqlController{}, "get:ShowMysql")

	beego.Router("/orm/", &controllers.OrmController{}, "get:ShowOrm")

	beego.Router("/ormInsert", &controllers.OrmController{}, "get:InsertOrm")

	beego.Router("/ormQuery", &controllers.OrmController{}, "get:QueryOrm")

	beego.Router("/ormUpdate", &controllers.OrmController{}, "get:UpdateOrm")

	beego.Router("/ormDelete", &controllers.OrmController{}, "get:DeleteOrm")

}
