package models

import (
	"github.com/astaxie/beego/orm"
	"myweb/myconf"
	_ "github.com/go-sql-driver/mysql"
)

//初始化
func init() {

	// 01 设置数据库的基本信息
	/**
	default : 数据库的别名
	mysql   : 数据名
	myconf.MySqlConfig["allconfig"] : 连接数据库 包括访问哪个数据库
	*/
	orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"])

	// 02 映射model 数据
	orm.RegisterModel(new(User), new(Acticle), new(ArticleType))

	// 03 生成表
	/**
	name 	数据库的别名
	force 	是否强制更新（如果表没有动的话为false）  ** 如果改成 true 会清空数据 重新建立表
	verbose: 是否可见（创建的时候）
	*/
	orm.RunSyncdb("default", false, true)

}

//表的设计  用户表   
// 一个用户可以看多篇文章
//多个用户也可以一个文章  
//多对多
type User struct {
	//注意这里的开头必须是大写  不然在数据库中不能创建
	Id       int
	Name     string     `orm:"unique"` //unique唯一的
	Pwd      string
	Articles []*Acticle `orm:"rel(m2m)"` //多对多
}
