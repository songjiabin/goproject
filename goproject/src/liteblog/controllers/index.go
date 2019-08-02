package controllers

import (
	"liteblog/models"
	"github.com/astaxie/beego/orm"

	"liteblog/syserror"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	//继承自它
	BaseController
}

// @router /  [get]
func (this *IndexController) Index() {

	//获取文章的数据
	//1、每页显示10条数据
	limit := 10
	//得到界面传过来的当前页码数据 没有就是1
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}

	title := this.GetString("title", "")

	notes, err := models.QueryNoteByPage(page, limit, title)
	if err != nil {
		this.Abort500(err)
	}
	this.Data["item"] = notes

	//得到总的数量
	count, err := models.QueryNoteCount(title)
	if err != nil {
		this.Abort500(err)
	}

	//根据总的条数 计算出要分多少页
	totpage := int(count) / limit
	if int(count)%limit != 0 {
		//取余数，不为0。那就要多加一页显示这些数据
		totpage = totpage + 1
	}

	// 将总页数 当前页 传到模版页面。等待渲染
	this.Data["totalpage"] = totpage
	this.Data["page"] = page

	this.TplName = "index.html"
}

// @router /about  [get]
func (this *IndexController) IndexAbout() {
	this.TplName = "about.html"
}

// @router /message  [get]
func (this *IndexController) IndexMessage() {
	this.TplName = "message.html"
}

// @router /login  [get]
func (this *IndexController) Login() {
	this.TplName = "user.html"
}

// @router /register  [get]
func (this *IndexController) Register() {
	this.TplName = "reg.html"
}

//获取文章的详情界面
// @router /details/:key [get]
func (this *IndexController) GetDetails() {
	//得到页面传过来的文章的key
	key := this.Ctx.Input.Param(":key")
	//根据key类进行查询文章
	logs.Info("key--------------------------->", key)
	note, err := models.QueryNoteByKey(key)
	if err != nil && err == orm.ErrNoRows {
		//没有查到
		this.Abort500(syserror.NewError("文章不存在", err))
	}

	//将文章信息给界面
	this.Data["note"] = note
	this.TplName = "details.html"
}

// @router /demo  [get]
func (this *IndexController) Demo() {
	this.TplName = "demo.html"
}
