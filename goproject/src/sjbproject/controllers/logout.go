package controllers

import "github.com/astaxie/beego"

type LogoutControllers struct {
	beego.Controller
}

//退出登录
func (this *LogoutControllers) Logout() {
	this.DelSession("userName")
	this.Redirect("/", 302)
}
