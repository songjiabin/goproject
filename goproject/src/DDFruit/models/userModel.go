package models

import "github.com/astaxie/beego/orm"

type User struct {
	BaseModel

	Id       int    `orm:"pk;auto"`         //id 自增 主键
	UserName string `orm:"size(20);unique"` //用户名 、唯一
	PassWord string `orm:"size(10)"`        //密码
	Email    string `orm:"size(50)"`        //邮箱
	Active   bool   `orm:"default(false)"`  //是否激活(只有邮箱认证通过了，才算是激活状态)  激活：true   否则 false
	Power    int    `orm:"default(0)"`      //权限设置   0表示未激活  1表示激活
}

// 新增用户 注册
func (this *User) InsertUser() (err error) {
	newOrm := orm.NewOrm()
	_, err = newOrm.Insert(this)
	if err != nil {
		
	} else {
		err = nil
	}
	return
}
