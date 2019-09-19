package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"sjbproject/models"
)

//查询所有的文章
type ShowArticleController struct {
	beego.Controller
}

//展示所有的文章类型
func (this *ShowArticleController) ShowArticle() {
	this.TplName = "indexOther.html"
	newOrm := orm.NewOrm()
	//查询文章列表的数据
	if queryArticleList(this, newOrm) {
		return
	}
	//查询文章的类型 并显示
	if queryArticleType(this, newOrm) {
		return
	}

}

//查询文章类型
func queryArticleType(this *ShowArticleController, newOrm orm.Ormer) (isError bool) {
	articleTypes := []models.ArticleType{}
	_, err := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		isError = true
		logs.Info("查询文章类型失败", err)
		return
	}
	this.Data["articleTypes"] = articleTypes
	return
}

//查询文章列表
func queryArticleList(this *ShowArticleController, newOrm orm.Ormer) (isError bool) {
	articles := []models.Article{}
	//查询所有的数目
	count, e := newOrm.QueryTable("Article").RelatedSel("ArticleType").Count()
	if e != nil {
		isError = true
		logs.Info("查询文章数量失败", e)
		return
	}
	totalPage = float64(count) / float64(pageSize)
	ceil := math.Ceil(totalPage) //向上取整  1.1 -- > 2
	//获取界面传来的pageIndex（要求跳转到哪一页）
	pageIndex, _ = this.GetInt("pageIndex")
	if float64(pageIndex) >= ceil {
		pageIndex = int(ceil)
	}
	if pageIndex == 0 {
		pageIndex = 1
	}

	//分页
	start := pageSize * (pageIndex - 1) //开始的位置
	newOrm.QueryTable("Article").RelatedSel("ArticleType").Limit(pageSize, start).All(&articles)

	//一共显示的页数
	this.Data["count"] = count
	//一共显示多少条记录

	this.Data["totalPage"] = int(ceil) //一共多少页
	this.Data["pageIndex"] = pageIndex //当前的页码数
	this.Data["articles"] = articles

}
