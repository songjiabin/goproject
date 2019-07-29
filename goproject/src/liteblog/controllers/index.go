package controllers

import "github.com/astaxie/beego/logs"

type IndexController struct {
	//继承自它
	BaseController
}

// @router /  [get]
func (this *IndexController) Index() {
	this.TplName = "index.html"
}

// @router /about  [get]
func (this *IndexController) IndexAbout() {
	this.TplName = "about.html"
}

// @router /message  [get]
func (this *IndexController) IndexMessage() {
	this.TplName = "message.html"
}



// @router /login  [get]
func (this *IndexController) Login() {
	logs.Info("出发我了。。。。IndexController")
	this.TplName = "user.html"
}
