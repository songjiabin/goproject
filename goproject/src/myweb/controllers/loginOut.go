package controllers

import "github.com/astaxie/beego"

type LoginOutController struct {
	beego.Controller
}

//退出登录
func (this *LoginOutController) LoginOut() {

	//删除session
	this.DelSession("userName")

	this.Redirect("/login", 302)
}
