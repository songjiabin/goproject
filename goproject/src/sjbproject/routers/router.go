package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"sjbproject/controllers"
)

func init() {

	//过滤router
	beego.InsertFilter("/article/*", beego.BeforeRouter, filtFun)

	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin")
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleReg")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.LogoutControllers{}, "get:Logout")

	beego.Router("/article/showArticle", &controllers.ArticleController{}, "get:ShowArticle;post:HandleArticle")
	beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/article/articleContent", &controllers.ArticleController{}, "get:ShowArticleContent")
	beego.Router("/article/deleteArticle", &controllers.ArticleController{}, "get:DeleteArticle")
	beego.Router("/article/updateArticle", &controllers.ArticleController{}, "get:UpdateArticle;post:HandleUpdateArticle")
	beego.Router("/article/addArticleType", &controllers.ArticleController{}, "get:ShowArticleType;post:HandleArticleType")
	beego.Router("/article/deleteArticleType", &controllers.ArticleController{}, "get:DeleteArticleType")
	beego.Router("/article/showArticleOther", &controllers.ShowArticleController{}, "get:ShowArticle")


	beego.Router("/layoutContent",&controllers.ShowLayoutContent{},"get:ShowLayoutContent")



}

//当匹配的时候 router之前进行的操作
var filtFun = func(ctx *context.Context) {
	session := ctx.Input.Session("userName")
	if session == nil {
		url := ctx.Request.URL
		logs.Info("请求的url事：", url)
		ctx.Redirect(302, "/")
	}
}
