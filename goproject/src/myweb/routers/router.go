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
	indexRouter()
	addArticleRouter()
	contentArticleRouter()
	updateArticleRouter()
	deleteArticleRouter()
	addArticleTypeRouter()
	loginOutRouter()
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
	beego.Router("/one2one", &controllers.OneToOneController{})
}

//注册的路由
func registerRouter() {
	beego.Router("/register", &controllers.RegisterController{})
}

//登录的路由
func loginRouter() {
	// get get方法
	// ShowLogin 路由中对应的方法
	// HandleLogin 路由
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
}

//index(首页的路由)
func indexRouter() {
	beego.Router("/index", &controllers.IndexController{}, "get:ShowIndex;post:SelectArtcileType")
}

// 添加文章
func addArticleRouter() {
	beego.Router("/addArticle", &controllers.AddController{}, "get:ShowAddArticle;post:AddArticle")
}

//文章内容
func contentArticleRouter() {
	beego.Router("/content", &controllers.ContentController{}, "get:ShowArticle")
}

//编辑
func updateArticleRouter() {
	beego.Router("/update", &controllers.UpdateControllerss{}, "get:UpdateArticle;post:HandlerUpdateArticle")
}

//删除
func deleteArticleRouter() {
	beego.Router("/delete", &controllers.DeleteArticleController{}, "get:DeleteActicle")
}

//添加文章的类型
func addArticleTypeRouter() {
	beego.Router("/addType", &controllers.AddTypeController{}, "get:AddType;post:HandleAddType")
}

func loginOutRouter()  {
	beego.Router("/loginout",&controllers.LoginOutController{},"get:LoginOut")
}
