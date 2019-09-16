package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

//展示文章信息
func (this *ArticleController) ShowArticle() {
	this.TplName = "index.html"
}

//处理文章列表
func (this *ArticleController) HandleArticle() {
	this.TplName = "index.html"
}

//展示添加文章列表的界面
func (this *ArticleController) ShowAddArticle() {
	this.TplName = "add.html"
}
