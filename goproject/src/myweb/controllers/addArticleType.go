package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
)

type AddTypeController struct {
	beego.Controller
}

//添加类型
func (c *AddTypeController) AddType() {

	//读取类型表显示数据
	newOrm := orm.NewOrm()
	var articleType []models.ArticleType

	_, e := newOrm.QueryTable("ArticleType").All(&articleType)
	if e != nil {
		beego.Info("查询类型错误", e)
		return
	}
	c.Data["addType"] = articleType

	c.TplName = "addType.html"
}

//处理添加类型的逻辑
func (c *AddTypeController) HandleAddType() {
	typeName := c.GetString("typeName")
	if typeName == "" {
		beego.Info("添加数据类型为空")
		c.Redirect("/addType", 302)
		return
	}
	//插入到表中
	newOrm := orm.NewOrm()
	articleType := models.ArticleType{TypeName: typeName}
	_, e := newOrm.Insert(&articleType)
	if e != nil {
		beego.Info("添加失败", e)
		c.Redirect("/addType", 302)
		return
	}

	c.Redirect("/addType", 302)
}
