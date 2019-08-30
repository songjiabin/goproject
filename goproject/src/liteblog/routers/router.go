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

	//给所有的note模块前面加上一个/note
	beego.AddNamespace(beego.NewNamespace("note", beego.NSInclude(&controllers.NoteController{})))



	beego.AddNamespace(beego.NewNamespace("message", beego.NSInclude(&controllers.MessageController{})))
}
