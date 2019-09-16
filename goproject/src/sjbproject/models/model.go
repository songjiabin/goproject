package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

//用户表
type User struct {
	Id       int
	UserName string `orm:"unique"`
	Passwd   string
}

//文章表
type Article struct {
	Id      int    `orm:"pk;auto"`                        //主键 自增
	Title   string `orm:"null;size(20)""`                 //标题
	Content string `orm:"size(200)"`                      //内容
	Img     string  `orm:"size(20)"`                                    // 图片
	Time    time.Time `orm:"type(datetime);auto_now_add"` //创建时间
	Count   int       `orm:"default(0)"`                  //阅读次数
	Num float64 `orm:"digits(12);decimals(4)"` //  浮点型 一共12位数，小数点后面4位
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", true, true)
}
