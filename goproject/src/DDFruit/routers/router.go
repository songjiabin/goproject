package routers

import (
	"DDFruit/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//用户的路由
	beego.Include(&controllers.UserControllers{})
	//错误处理的界面
	beego.ErrorController(&controllers.ErrorController{})
}

