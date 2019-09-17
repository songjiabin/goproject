package routers

import (
	"github.com/astaxie/beego"
	"sjbproject/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin")
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleReg")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/showArticle", &controllers.ArticleController{}, "get:ShowArticle;post:HandleArticle")
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")

}
