package models

import "time"

// 文章的结构体
type Acticle struct {
	Id       int   `orm:"pk;auto"`  					// pk 主键  auto 自增
	Aname    string  `orm:"default(默认文章名字)"'`      //文章名称
	Atime    time.Time `orm:"auto_now"`								   //文章的时间
	Acount   int      `orm:"default(0);null"`   	   //阅读量     null 允许为空
	Acontent string    `orm:"size(100)"`			  //文章的内容
	Aimg     string    								  //图片 存放图片的路径
}
