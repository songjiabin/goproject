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
	Article  []*Article `orm:"reverse(many)"` //多对多
}

//文章表
type Article struct {
	Id          int          `orm:"pk;auto"`                          //主键 自增
	Title       string       `orm:"size(20)""`                        //标题
	Content     string       `orm:"size(200)"`                        //内容
	Img         string       `orm:"size(100)"`                        // 图片
	Time        time.Time    `orm:"type(datetime);auto_now_add;null"` //创建时间
	Count       int          `orm:"default(0);null"`                  //阅读次数
	Num         float64      `orm:"digits(12);decimals(4);null"`      //  浮点型 一共12位数，小数点后面4位
	ArticleType *ArticleType `orm:"rel(fk)"`                          //外键  设置一对多的关系
	User        []*User      `orm:"rel(m2m)"`                         ////多对多
}

//类型表   一个文章只属于一个类型 ，但是一个类型中有多个文章
type ArticleType struct {
	Id       int        `orm:"pk;auto"`
	TypeName string     `orm:"size(20);unique"`
	Article  []*Article `orm:"reverse(many)"` //一对多的反向关系
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}
