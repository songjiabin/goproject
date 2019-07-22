package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"math"
)

type IndexController struct {
	beego.Controller
}

//每页多少条数据
const pageSize = 3

//展示index
func (c *IndexController) ShowIndex() {

	//查询内容
	newOrm := orm.NewOrm()

	var acticles []models.Acticle

	//查询表
	_, e := newOrm.QueryTable("Acticle").All(&acticles)
	if e != nil {
		beego.Info("查询数据失败了")
		return
	}
	c.Data["articles"] = acticles

	//查询数据的条数
	count, queryError := newOrm.QueryTable("Acticle").Count()
	if queryError != nil {
		beego.Info("查询数量错误", queryError)
		return
	}

	//总共多少条记录
	c.Data["count"] = count
	//一共多少页
	totalPage := float64(count) / float64(pageSize)
	c.Data["totalPage"] = math.Ceil(totalPage)

	c.TplName = "index.html"
}
