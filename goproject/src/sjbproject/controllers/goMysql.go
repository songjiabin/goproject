package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlController struct {
	beego.Controller
}

type SqlBean struct {
	id   int
	name string
}

func (c *MySqlController) ShowMysql() {

	c.TplName = "mysql.html"

	// mysql -u root -p
	// 使用 tcp 协议
	// mydb使用的数据库
	db, e := sql.Open("mysql", "root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8")
	if e != nil {
		beego.Info("连接错误", e)
		c.Data["data"] = "连接错误"
		return
	}
	defer db.Close()
	//操作数据库
	_, e = db.Exec("create table  if  not exists   useInfo(id int,name varchar(10)) ")
	if e != nil {
		beego.Info("创建表错误", e)
		c.Data["data"] = "创建表错误"
		return
	}

	beego.Info("创建表成功")
	c.Data["data"] = "创建表成功"
	c.Ctx.WriteString("创建表成了！！！")

	//插入数据
	/*_, e = db.Exec("insert into  useInfo  value (?,?)", 1, "宋佳宾")
	if e != nil {
		beego.Info("插入数据错误", e)
		return
	}

	c.Ctx.WriteString("插入数据成功了...")*/

	result, e := db.Query("select id from useInfo where id =?", 1)
	if e != nil {
		beego.Info("查询失败", e)
		return
	}
	//beanList := make([]SqlBean,0)
	var id string
	for result.Next() {
		result.Scan(&id)
		beego.Info("查询成功 id-->", id)
	}



}
