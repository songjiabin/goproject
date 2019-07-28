package models

import (
	"github.com/astaxie/beego/orm"
	"myweb/myconf"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdll := 30 //设置最大空闲连接
	maxConn := 30 //设置最大的连接数
	//orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"])
	orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"], maxIdll, maxConn)
	// 02 映射model 数据
	orm.RegisterModel()

}
