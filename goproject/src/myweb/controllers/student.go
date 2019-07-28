package controllers

import "github.com/astaxie/beego"

type StudentControllers struct {
	beego.Controller
}


//添加学生
func (this *StudentControllers) AddStudent() {
	this.TplName = "addStudent.html"
	this.Layout = "layout.html"
}
