package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"path"
	"time"
)

type AddController struct {
	beego.Controller
}

//添加文章的展示界面
func (c *AddController) ShowAddArticle() {

	//进入添加文章的界面

	//展示所有的文章类型
	newOrm := orm.NewOrm()
	var arctcileType []models.ArticleType
	_, errOfArticleType := newOrm.QueryTable("ArticleType").All(&arctcileType)
	if errOfArticleType != nil {
		beego.Info("获取文章类型错误", errOfArticleType)
		return
	}

	c.Data["arctcileType"] = arctcileType

	c.TplName = "add.html"

}

// 添加文章的post请求
func (c *AddController) AddArticle() {

	c.TplName = "add.html"

	//文章的标题
	acticleName := c.GetString("articleName")

	//文章的内容
	articleContent := c.GetString("content")

	//拿到图片
	file, header, e := c.GetFile("uploadname")
	defer file.Close()

	//1、对图片限定格式
	//得到文件的后缀
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" {
		beego.Info("上传文件的格式错误", ext)
		return
	}

	//2、限制大小
	if header.Size > 50000000 {
		beego.Info("上传的文件太大了！！！")
		return
	}

	//3、需要对文件重命名，防止文件重复
	//用时间进行重命名
	fileName := time.Now().Format("2006-01-02-13-04-05")

	fmt.Println("时间是：", fileName)

	if e != nil {
		beego.Info("上传图片失败了...")
		fmt.Println("get filr error:", e)
		return
	} else {
		//保存文件
		c.SaveToFile("uploadname", "./static/img/"+fileName+ext)
		beego.Info(fileName + ext)
	}

	//获取文章的类型
	typeName := c.GetString("select")
	if typeName == "" {
		beego.Info("文章类型不能为空！")
		return
	}

	//根据名字查到 ArticleType 对象
	var articleType models.ArticleType

	articleTypeOrm := orm.NewOrm()
	articleType.TypeName = typeName
	articleTypeOrm.Read(&articleType, "TypeName")

	//插入数据到数据库中去
	newOrm := orm.NewOrm()

	acticle := models.Acticle{}
	acticle.Aname = acticleName
	acticle.Acontent = articleContent
	acticle.Aimg = "/static/img/" + fileName + ext
	//插入类型
	acticle.AType = &articleType
	_, e2 :=
		newOrm.Insert(&acticle)
	if e2 != nil {
		beego.Info("插入数据出错了哦...", e2)
		return
	}

	//否则则跳转到index界面中去
	c.Redirect("/index", 302)

}
