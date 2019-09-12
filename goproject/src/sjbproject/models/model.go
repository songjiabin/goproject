package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id   int
	Name string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
