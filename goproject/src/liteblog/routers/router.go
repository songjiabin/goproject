package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {

	//注册路由
	beego.Include(&controllers.IndexController{})

}
