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
const pageSize = 2

//当前的页数
var pageIndex = 1

//展示index
func (c *IndexController) ShowIndex() {

	//查询内容
	newOrm := orm.NewOrm()

	var acticles []models.Acticle

	//查询数据的条数
	count, queryError := newOrm.QueryTable("Acticle").RelatedSel("ArticleType").Count()
	if queryError != nil {
		beego.Info("查询数量错误", queryError)
		return
	}

	pageIndex, es := strconv.Atoi(c.GetString("pageIndex"))
	if es != nil {
		pageIndex = 1
	}

	start := pageSize * (pageIndex - 1) //开始的位置
	//Limit 截取
	//pageSize 一页显示多少
	//start 从哪里开始查
	_, e := newOrm.QueryTable("Acticle").Limit(pageSize, start).RelatedSel("ArticleType").All(&acticles)

	if e != nil {
		beego.Info("查询数据失败了")
		return
	}

	isFirstPage := false //是否为第一页
	//当pageindex ==1 的时候 上一页默认不让点击
	if pageIndex == 1 {
		isFirstPage = true
	}

	//总共多少条记录
	c.Data["count"] = count
	//一共多少页
	totalPage := float64(count) / float64(pageSize)
	c.Data["totalPage"] = math.Ceil(totalPage)
	c.Data["pageIndex"] = pageIndex
	c.Data["isFirstPage"] = isFirstPage
	c.Data["articles"] = acticles

	isEndPage := false
	if float64(pageIndex) == math.Ceil(totalPage) {
		isEndPage = true
	}

	c.Data["isEndPage"] = isEndPage

	//获取数据类型
	var arctcileType []models.ArticleType
	_, errOfArticleType := newOrm.QueryTable("ArticleType").All(&arctcileType)
	if errOfArticleType != nil {
		beego.Info("获取文章类型错误", errOfArticleType)
		return
	}

	c.Data["arctcileType"] = arctcileType

	c.TplName = "index.html"
}

//选择下拉框
func (c *IndexController) SelectArtcileType() {
	typeName := c.GetString("select")
	beego.Info(typeName)
	if typeName == "" {
		beego.Info("数据传递失败")
		return
	}

	//得到了下拉菜单的数据
	newOrm := orm.NewOrm()
	var acticles [] models.Acticle
	//因为是多表联查
	// Acticle表中存在ArticleType的外键
	//所以ArticleType__表示去ArticleType表中去查
	//RelatedSel 建立连接
	_, e := newOrm.QueryTable("Acticle").RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&acticles)
	if e != nil {
		beego.Info("查询错误", e)
		return
	}
	beego.Info("根据下拉菜单选择数据：-》", acticles)
	//c.Redirect("/index", 302)
}
