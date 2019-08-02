package models

import (
	"github.com/astaxie/beego/orm"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

type Note struct {
	Id      int
	Key     string `orm:"unique"`
	User    *User  `orm:"rel(fk)"` //当前的文章属于哪个用户
	Title   string `orm:"size(200)"`
	Summary string `orm:"size(800)"`             //摘要
	Content string `orm:"size(2000)"`            //内容
	Visist  int    `orm:"size(2000);default(0)"` //浏览次数
	Praise  int    `orm:"size(2000)"`            //点赞的次数
}

//根据 key 查询是否有这个数据
func QueryNoteByKey(key string) (note Note, err error) {
	newOrm := orm.NewOrm()
	note = Note{
		Key: key,
	}
	err = newOrm.Read(&note, "Key")
	return
}

func InsertNote(note *Note) (err error) {
	newOrm := orm.NewOrm()
	_, e := newOrm.Insert(note)
	err = e
	return
}

//获取摘要
func GetSummary(html string) (string, error) {
	var bufbytes bytes.Buffer
	_, err := bufbytes.Write([]byte(html))
	if err != nil {
		return "", err
	}
	document, err := goquery.NewDocumentFromReader(&bufbytes)
	if err != nil {
		return "", err
	}

	htmlStr := document.Find("body").Text()
	//如果大于600，那么就取前600个字符串
	if len(htmlStr) > 600 {
		htmlStr = htmlStr[:600]
	}
	return htmlStr, nil

}

// 根据当前页数和每页显示的行数得到文章数据的方法
// page 	当前的页数
// limit	每页显示的数量
// title    根据筛选内容查
func QueryNoteByPage(page, limit int, title string) (notes []Note, err error) {
	newOrm := orm.NewOrm()
	notes = []Note{}
	if title == "" {
		_, err = newOrm.QueryTable("Note").Limit(limit, (page-1)*limit).All(&notes)
	} else {
		_, err = newOrm.QueryTable("Note").Filter("Title__contains", title).Limit(limit, (page-1)*limit).All(&notes)
	}

	logs.Info("结果是：", notes)

	return
}

//得到总的数量
func QueryNoteCount(title string) (count int64, err error) {
	newOrm := orm.NewOrm()

	if title == "" {
		count, err = newOrm.QueryTable("Note").Count()
	} else {
		count, err = newOrm.QueryTable("Note").Filter("Title__contains", title).Count()
	}

	return
}

//根据文章的key 和 用户的id 查询文章
func QueryNoteByUserIdAndNoteKey(id int, key string) (note Note, err error) {
	newOrm := orm.NewOrm()
	err = newOrm.QueryTable("Note").RelatedSel("User").Filter("Key", key).Filter("User__Id", id).One(&note)
	return
}

//更新
func UpdateNoteByKey(note Note) (err error) {
	newOrm := orm.NewOrm()
	_, e := newOrm.Update(&note)
	return e
}

//删除
func DeleteNoteByKey(key string) error {
	newOrm := orm.NewOrm()
	note := Note{Key: key}
	i, e := newOrm.Delete(&note,"Key")
	logs.Info("删除了。。。。。",i,e)
	return e
}
