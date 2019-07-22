package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
)

//删除

type DeleteArticleController struct {
	beego.Controller
}

//删除
func (c *DeleteArticleController) DeleteActicle() {
	id, e := c.GetInt("id")
	if e != nil {
		beego.Info(e)
		return
	}
	//根据id删除数据库中数据
	article := models.Acticle{Id: id}
	newOrm := orm.NewOrm()
	_, e2 := newOrm.Delete(&article)
	if e2 != nil {
		beego.Info(e2)
		return
	} else {
		beego.Info("删除完成!")
	}

	c.Redirect("/index", 302)

}
