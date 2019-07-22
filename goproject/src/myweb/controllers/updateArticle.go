package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"time"
	"path"
)

// 修改文章的信息
type UpdateControllerss struct {
	beego.Controller
}

func (c *UpdateControllerss) UpdateArticle() {

	id, e := c.GetInt("id")
	if e != nil {
		return
	}

	//根据id查询数据  并显示
	newOrm := orm.NewOrm()
	article := models.Acticle{Id: id}
	newOrm.Read(&article)
	c.Data["article"] = article
	c.TplName = "update.html"
}

//处理更新数据的方法
func (c *UpdateControllerss) HandlerUpdateArticle() {
	id, eGetId := c.GetInt("id")
	if eGetId != nil {
		beego.Info(eGetId)
		return
	}
	//文章标题
	arcticleName := c.GetString("articleName")
	//文章内容
	arcticleContent := c.GetString("content")
	//文章的图片
	file, header, eFile := c.GetFile("uploadname")
	defer file.Close()

	fileName := time.Now().Format("2006-01-02-13-04-05")
	ext := path.Ext(header.Filename)

	if eFile != nil {
		beego.Info(eFile)
		return
	} else {
		c.SaveToFile("uploadname", "./static/img/"+fileName+ext)
	}

	arctile := models.Acticle{Id: id}
	arctile.Aname = arcticleName
	arctile.Acontent = arcticleContent
	arctile.Aimg = "/static/img/" + fileName + ext
	newOrm := orm.NewOrm()
	_, eUpdate := newOrm.Update(&arctile)
	if eUpdate != nil {
		beego.Info(eUpdate)
		return
	} else {
		c.Redirect("/index", 302)
	}

	c.TplName = "update.html"

}
