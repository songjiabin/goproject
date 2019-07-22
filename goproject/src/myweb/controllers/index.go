package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"math"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

//每页多少条数据
const pageSize = 1

//当前的页数
var pageIndex = 1

//展示index
func (c *IndexController) ShowIndex() {

	//查询内容
	newOrm := orm.NewOrm()

	var acticles []models.Acticle

	//查询数据的条数
	count, queryError := newOrm.QueryTable("Acticle").Count()
	if queryError != nil {
		beego.Info("查询数量错误", queryError)
		return
	}

	pageIndex, _ = strconv.Atoi(c.GetString("pageIndex"))

	start := pageSize * (pageIndex - 1) //开始的位置
	//Limit 截取
	//pageSize 一页显示多少
	//start 从哪里开始查
	_, e := newOrm.QueryTable("Acticle").Limit(pageSize, start).All(&acticles)

	if e != nil {
		beego.Info("查询数据失败了")
		return
	}

	//总共多少条记录
	c.Data["count"] = count
	//一共多少页
	totalPage := float64(count) / float64(pageSize)
	c.Data["totalPage"] = math.Ceil(totalPage)
	c.Data["pageIndex"] = pageIndex

	c.Data["articles"] = acticles

	c.TplName = "index.html"
}
