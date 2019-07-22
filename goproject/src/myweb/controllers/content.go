package controllers

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"myweb/models"
)

//查看详情的控制器
type ContentController struct {
	beego.Controller
}

//展示文章内容
func (c *ContentController) ShowArticle() {

	id, e := c.GetInt("id")
	beego.Info("错误是：", e)

	//根据id获取数据库数据
	newOrm := orm.NewOrm()
	article := models.Acticle{Id: id}
	err := newOrm.Read(&article)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}

	//将数据给视图
	c.Data["article"] = article

	c.TplName = "content.html"

}
