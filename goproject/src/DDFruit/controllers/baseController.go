package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//所有的都会执行
func (this *BaseController) Prepare() {

}

//500错误的时候
func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}
