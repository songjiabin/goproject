package controllers

import (
	"bytes"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"io"
	"math"
	"sjbproject/models"
)

//查询所有的文章
type ShowArticleController struct {
	beego.Controller
}

//必执行的操作
func (this *ShowArticleController) Prepare() {
	logs.Info("Prepare------------------>>>>>>")
}

//展示所有的文章类型
func (this *ShowArticleController) ShowArticle() {
	this.TplName = "indexOther.html"
	newOrm := orm.NewOrm()
	//查询文章的类型 并显示
	if error := queryArticleType(this, newOrm); error {
		return
	}
	//查询文章列表的数据
	if queryArticleList(this, newOrm) {
		return
	}

}

//查询文章类型
func queryArticleType(this *ShowArticleController, newOrm orm.Ormer) (isError bool) {
	articleTypes := []models.ArticleType{}

	//获取连接池
	rc := models.RedisClient.Get()
	defer rc.Close()

	//读取到的字节数组格式
	b, err := redis.Bytes(rc.Do("get", "article_type"))
	if err != nil {
		logs.Info("读取redis失败", err)
	}
	//将读取到的数据进行转换成实体
	decoder := gob.NewDecoder(bytes.NewReader(b)) //解码器
	err = decoder.Decode(&articleTypes)
	if err != nil && err == io.EOF {
		logs.Info("从数据库中读取数据.....")
		_, err := newOrm.QueryTable("ArticleType").All(&articleTypes)
		if err != nil {
			isError = true
			logs.Info("查询文章类型失败", err)
			return
		}
		//定义存储数据的容器
		var buffer bytes.Buffer

		encoder := gob.NewEncoder(&buffer) //进行编码
		encoder.Encode(articleTypes)
		//将字节数据存储到redis 中
		_, err = rc.Do("set", "article_type", buffer.Bytes())
		if err != nil {
			logs.Info("redis数据库操作错误", err)
			return
		}

	} else if err != nil {
		logs.Info("解析字节出错", err)
	}
	logs.Info(articleTypes)

	//如果没有从redis中查询到数据，那么从数据库中查询

	//赋值
	this.Data["articleTypes"] = articleTypes

	return
}

//查询文章列表
func queryArticleList(this *ShowArticleController, newOrm orm.Ormer) (isError bool) {
	articles := []models.Article{}
	var count int64 = 0
	var e error
	//是否有查询条件(是否有下拉框）
	selectName := this.GetString("select")

	//当点击下一页的时候带过来的文章类型
	currentTypeName := this.GetString("currentTypeName")

	if currentTypeName != "" {
		selectName = currentTypeName
	}

	if selectName == "" {
		//如果是空，表明没有筛选条件。直接返回所有的数据
		//查询所有的数目
		//count, e = newOrm.QueryTable("Article").RelatedSel("ArticleType").Count()
		articleType := models.ArticleType{}
		e := newOrm.QueryTable("ArticleType").One(&articleType)
		if e != nil {
			isError = true
			logs.Info("查询文章类型失败", e)
			return
		}
		selectName = articleType.TypeName
	}
	//否则根据筛选条件进行查找数据
	count, e = newOrm.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", selectName).Count()

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

	if selectName == "" {
		//如果是空，表明没有筛选条件。直接返回所有的数据
		//查询所有的数目
		newOrm.QueryTable("Article").RelatedSel("ArticleType").Limit(pageSize, start).All(&articles)
	} else {
		//否则根据筛选条件进行查找数据
		newOrm.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", selectName).
			Limit(pageSize, start).All(&articles)
	}

	//一共显示的页数
	this.Data["count"] = count
	//一共显示多少条记录

	this.Data["totalPage"] = int(ceil) //一共多少页
	this.Data["pageIndex"] = pageIndex //当前的页码数
	this.Data["articles"] = articles
	this.Data["currentTypeName"] = selectName //当前选择的名字
	return
}
