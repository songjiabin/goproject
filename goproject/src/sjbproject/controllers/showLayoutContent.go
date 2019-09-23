package controllers

import "github.com/astaxie/beego"

type ShowLayoutContent struct {
	beego.Controller
}

func (this *ShowLayoutContent) ShowLayoutContent() {
	this.Layout = "layout.html"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["title_text"] = "title_text.html"
	this.Data["title"]="首页"
	this.TplName = "main.html"
}
