package routers

import (
	"myweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//路由 进行匹配 分发
	demoRouter()
	registerRouter()
	loginRouter()
}

//案例的小demo 增删改查
func demoRouter() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/songjiabin", &controllers.TestController{})
	beego.Router("/post", &controllers.PostController{})
	beego.Router("/insertUser", &controllers.OrmController{})
	beego.Router("/selectUser", &controllers.OrmSelectController{})
	beego.Router("/updateUser", &controllers.UpdateController{})
	beego.Router("/deleteUser", &controllers.DeleteController{})
}

//注册的路由
func registerRouter() {
	beego.Router("/register", &controllers.RegisterController{})
}

func loginRouter() {
	// get get方法
	// ShowLogin 路由中对应的方法
	// HandleLogin 路由
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
}
