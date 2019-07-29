package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//注册错误的路由
	beego.ErrorController(&controllers.ErrorController{})
	//主页路由
	beego.Include(&controllers.IndexController{})
	//登录路由
	beego.Include(&controllers.UserController{})

}
