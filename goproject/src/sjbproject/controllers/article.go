package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"path"
	"sjbproject/models"
	"sjbproject/utils"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//展示文章信息
func (this *ArticleController) ShowArticle() {
	this.TplName = "index.html"
	articles := []models.Article{};
	newOrm := orm.NewOrm()
	newOrm.QueryTable("article").All(&articles)
	this.Data["articles"] = articles
}

//处理文章列表
func (this *ArticleController) HandleArticle() {
	this.TplName = "index.html"
}

//展示添加文章列表的界面
func (this *ArticleController) ShowAddArticle() {
	this.TplName = "add.html"
}

//处理添加文章
func (this *ArticleController) HandleAddArticle() {
	this.TplName = "add.html"
	articleName := this.GetString("articleName")
	content := this.GetString("content")

	//获取文件
	file, header, e := this.GetFile("uploadname")
	defer file.Close()
	if e != nil {
		logs.Debug("得到文件失败", e)
		return
	}
	//获取文件的格式类型
	ext := path.Ext(header.Filename)
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		logs.Debug("上传图片的格式有问题！")
		return
	}

	//判断文件的大小
	if header.Size > 1024*100 {
		logs.Debug("文件太大了，不允许上传")
		return
	}
	//将图片的名字修改下  不能重名
	format := time.Now().Format("2006-01-02-15-04-05")

	//保存图片
	imgPath := "./static/img/" + format + ext
	e = this.SaveToFile("uploadname", imgPath)
	if e != nil {
		logs.Debug("保存文件失败", e)
		return
	}

	//插入数据
	newOrm := orm.NewOrm()
	article := models.Article{Content: content, Title: articleName, Img: imgPath}
	_, e = newOrm.Insert(&article)
	logs.Debug(len(imgPath))
	utils.HandleError("插入article错误", e)

	logs.Debug("插入成功")
	this.Redirect("/showArticle", 302)

}
