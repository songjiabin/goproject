package routers

import (
	"github.com/astaxie/beego"
	"myweb/controllers"
)

func init() {
	beego.Router("/addStudent", &controllers.StudentControllers{},"get:AddStudent")
}
