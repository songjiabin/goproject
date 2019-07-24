package models

import "time"

// 文章的结构体  文章表和文章类型表 一对多
//一个文章对应一个类型
//一个类型下面有多个文章
type Acticle struct {
	Id          int          `orm:"pk;auto"`          // pk 主键  auto 自增
	Aname       string       `orm:"default(默认文章名字)"'` //文章名称
	Atime       time.Time    `orm:"auto_now"`         //文章的时间
	Acount      int          `orm:"default(0);null"`  //阅读量     null 允许为空
	Acontent    string       `orm:"size(100)"`        //文章的内容
	Aimg        string                                //图片 存放图片的路径
	ArticleType *ArticleType `orm:"rel(fk)"`          //文章的类型 rel对应 类型表的reverse
	AUsers      [] *User     `orm:"reverse(many)"`    //一个文章可以被多个用户看  reverse(many) 这在前面的话  创建表的名字为：user_acticles
}
