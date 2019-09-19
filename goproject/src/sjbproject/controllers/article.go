package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"sjbproject/models"
	"sjbproject/utils"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

const pageSize int = 1 //每个页多少条数据
var pageIndex = 1       //当前的页数
var totalPage float64   //一共要显示的页数

//展示文章信息
func (this *ArticleController) ShowArticle() {
	this.TplName = "index.html"

	articles := []models.Article{};
	newOrm := orm.NewOrm()


	se := this.GetString("select")
	if se == "" {
		logs.Info("下拉框获取数据失败")
		return
	}


	//查询所有的数目
	count, _ := newOrm.QueryTable("Article").RelatedSel("ArticleType").Count()
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



	//查询所有的文章类型
	articleTypes := []models.ArticleType{}
	_, e := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if e != nil {
		logs.Info("查询文章类型失败...")
		return
	}
	this.Data["articleTypes"] = articleTypes

}

//处理文章列表 从index.html发送过来的数据
func (this *ArticleController) HandleArticle() {

	this.TplName = "index.html"

	se := this.GetString("select")
	if se == "" {
		logs.Info("下拉框获取数据失败")
		return
	}

	articles := []models.Article{}
	newOrm := orm.NewOrm()

	//查询所有的数目
	count, _ := newOrm.QueryTable("Article").RelatedSel("ArticleType").Count()
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
	//多表联查 	//更新文章的类型来筛选文章
	newOrm.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", se).
		Limit(pageSize, start).All(&articles)

	//一共显示的页数
	this.Data["count"] = count
	//一共显示多少条记录

	this.Data["totalPage"] = int(ceil) //一共多少页
	this.Data["pageIndex"] = pageIndex //当前的页码数
	this.Data["articles"] = articles

	//查询所有的文章类型
	articleTypes := []models.ArticleType{}
	_, e := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if e != nil {
		logs.Info("查询文章类型失败...")
		return
	}
	this.Data["articleTypes"] = articleTypes
	this.Data["CurrentTypeName"] = se
}

//展示添加文章列表的界面
func (this *ArticleController) ShowAddArticle() {
	//查询所有的文章类型
	articleTypes := []models.ArticleType{}
	newOrm := orm.NewOrm()
	_, e := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if e != nil {
		logs.Info("查询文章类型失败...")
		return
	}
	this.Data["articleTypes"] = articleTypes
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

	newOrm := orm.NewOrm()

	//获取文件的类型
	sel := this.GetString("select")
	if sel == "" {
		logs.Info("获取文章类型为空")
		return
	}
	articleType := models.ArticleType{TypeName: sel}
	e = newOrm.Read(&articleType, "TypeName")
	if e != nil {
		logs.Info("查询对应的文章类型失败", e)
		return
	}

	//插入数据
	article := models.Article{Content: content, Title: articleName, Img: imgPath}
	article.ArticleType = &articleType
	_, e = newOrm.Insert(&article)
	logs.Debug(len(imgPath))
	utils.HandleError("插入article错误", e)

	logs.Debug("插入成功")
	this.Redirect("/showArticle", 302)

}

//展示文章详情界面
func (this *ArticleController) ShowArticleContent() {
	this.TplName = "content.html"
	id := this.GetString("id")
	//查询数据库对应的数据
	newOrm := orm.NewOrm()
	idInt, _ := strconv.Atoi(id)
	article := models.Article{}
	//e := newOrm.Read(&article, "Id")
	e := newOrm.QueryTable("Article").Filter("Id", idInt).RelatedSel("ArticleType").One(&article)
	if e != nil {
		logs.Debug("查询文章详情错误", e)
		return
	}
	article.Count += 1
	newOrm.Update(&article)
	this.Data["article"] = article
}

//删除文章
func (this *ArticleController) DeleteArticle() {
	this.TplName = "index.html"
	id := this.GetString("id")
	newOrm := orm.NewOrm()
	idInt, _ := strconv.Atoi(id)
	article := models.Article{Id: idInt}
	_, e := newOrm.Delete(&article, "Id")
	if e != nil {
		logs.Debug("删除失败", e)
		return
	}
	this.Redirect("/showArticle", 302)
}

//更新文章
func (this *ArticleController) UpdateArticle() {
	this.TplName = "update.html"
	id := this.GetString("id")
	newOrm := orm.NewOrm()
	idInt, _ := strconv.Atoi(id)
	article := models.Article{Id: idInt}
	e := newOrm.Read(&article, "Id")
	if e != nil {
		logs.Debug("查询文章失败", e)
		return
	}
	this.Data["article"] = article

}

//处理更新
func (this *ArticleController) HandleUpdateArticle() {
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

	//获取需要更新的id
	id := this.GetString("id")
	idInt, _ := strconv.Atoi(id)

	article := models.Article{Id: idInt}
	newOrm := orm.NewOrm()

	//查询文章是否存在
	e = newOrm.Read(&article, "Id")
	if e != nil {
		logs.Debug("要更新的文章不存在", e)
		return
	}

	//给赋值操作
	article.Title = articleName
	article.Content = content
	article.Img = imgPath

	_, e = newOrm.Update(&article)
	if e != nil {
		logs.Debug("更新文章失败", e)
		return
	}
	logs.Debug("更新文章成功")
	this.Redirect("/showArticle", 302)
}

//处理文章类型=
func (this *ArticleController) ShowArticleType() {
	this.TplName = "addType.html"
	articleTypes := []models.ArticleType{}
	newOrm := orm.NewOrm()
	_, e := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if e != nil {
		logs.Info("查询文章类型失败", e)
		return
	}
	this.Data["articleTypes"] = articleTypes
}

//处理添加类型
func (this *ArticleController) HandleArticleType() {
	typeName := this.GetString("typeName")
	if typeName == "" {
		logs.Info("插入文章类型不能为空")
		this.TplName = "addType.html"
		return
	}
	//上传到服务器
	articleType := models.ArticleType{TypeName: typeName}
	newOrm := orm.NewOrm()
	_, e := newOrm.Insert(&articleType)
	if e != nil {
		logs.Info("插入文章类型失败", e)
		return
	}
	this.Redirect("/addArticleType", 302)
}

//删除类型
func (this *ArticleController) DeleteArticleType() {
	id, _ := this.GetInt("id")
	articleType := models.ArticleType{Id: id}
	newOrm := orm.NewOrm()
	_, e := newOrm.Delete(&articleType, "Id")
	if e != nil {
		logs.Info("删除类型错误", e)
		return
	}
	this.Redirect("/addArticleType", 302)
}
