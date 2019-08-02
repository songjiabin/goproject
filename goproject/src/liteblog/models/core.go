package models

import (
	"github.com/astaxie/beego/orm"
	"liteblog/myconf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/logs"


)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdll := 30 //设置最大空闲连接
	maxConn := 30 //设置最大的连接数
	//orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"])
	orm.RegisterDataBase("default", "mysql", myconf.MySqlConfig["allconfig"], maxIdll, maxConn)
	// 02 映射model 数据
	orm.RegisterModel(new(User), new(Note))
	orm.RunSyncdb("default", false, true)

	//在程序启动的时候插入admin用户

	user := User{
		Id:     1,
		Name:   "admin",
		Email:  "admin@qq.com",
		Pwd:    "123456",
		Avatar: "/static/images/info-img.png",
		Role:   0, //管理员的身份
	}
	newOrm := orm.NewOrm()
	//查询是否存在
	error := newOrm.Read(&user)
	//查询不到此列的时候
	if error != nil && error == orm.ErrNoRows {
		logs.Info("错误shi->：", error)
		newOrm.Insert(&user)
		return
	} else {
		logs.Info("此用户已经存在")
	}

}

