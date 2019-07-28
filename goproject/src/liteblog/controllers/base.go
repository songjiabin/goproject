package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//所有的路由都会继承此方法
//首先执行的方法
func (this *BaseController) Prepare() {
	//请求的url
	this.Data["path"] = this.Ctx.Request.RequestURI
}
