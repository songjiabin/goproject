package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {

}

func init() {
	logs.Info("-------------------------------------------")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdll := 30 //设置最大空闲连接
	maxConn := 30 //设置最大的连接数
	orm.RegisterDataBase("default", "mysql", "root:123root@A@tcp(192.144.238.85:3306)/ddFruit?charset=utf8&loc=Asia%2FShanghai", maxIdll, maxConn)
	// 02 映射model 数据
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}