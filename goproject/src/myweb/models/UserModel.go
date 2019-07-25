package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myweb/myconf"
)

//初始化
func init() {

	//**使用原生的方法创建表**

	useMySelf()

	//告诉系统,我要操作的数据库是什么数据库
	// mysql sqllite3 等默认注册好了 不需要额为注册
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 01 设置数据库的基本信息
	/**
	default : 数据库的别名（必须注册一个别名为default的数据库，作为默认使用） 用于切换数据库的时候使用
	mysql   : 数据名
	myconf.MySqlConfig["allconfig"] : 连接数据库 包括访问哪个数据库
	*/
	maxIdll := 30 //设置最大空闲连接
	maxConn := 30 //设置最大的连接数
	//orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"])
	orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"], maxIdll, maxConn)

	// 02 映射model 数据
	orm.RegisterModel(new(User), new(Acticle), new(ArticleType), new(PersonName), new(Person))

	// 03 生成表
	/**
	name 	数据库的别名
	force 	是否强制更新（如果表没有动的话为false）  ** 如果改成 true 会清空数据 重新建立表
	verbose: 是否可见(显示编译好的sql语句)（创建的时候）
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
	Name     string `orm:"unique"` //unique唯一的
	Pwd      string
	Articles []*Acticle `orm:"rel(m2m)"` //多对多
}

//用自己的方式进行创建表
func useMySelf() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test_song")
	if err != nil {
		fmt.Print("打开数据库失败！")
		return
	}
	defer db.Close()

	sql := `create table if not exists beego_test(
  		id INT PRIMARY KEY auto_increment,  
 		 name VARCHAR(20))`

	//执行创建表工作
	_, errs := db.Exec(sql)
	if errs != nil {
		fmt.Print("创建表失败", errs)
		return
	}

	fmt.Print("创建表成功了")

	sql_insert := `insert into  beego_test(name) values ("许美琳")`
	_, err3 := db.Exec(sql_insert)
	if (err3 != nil) {
		fmt.Print("插入失败")
		return
	}
	fmt.Print("插入成功")

	sql_select := `select * from beego_test 
where (
 id=1
)`
	result, err4 := db.Exec(sql_select)
	if err4 != nil {
		fmt.Print("查询失败", err4)
		return
	}

	fmt.Print(result)
}
