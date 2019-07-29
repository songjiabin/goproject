package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int    `orm:"pk;auto"`
	Name   string `orm:"unique"`
	Email  string `orm:"unique"`
	Pwd    string
	Avatar string                    //头像
	Role   int    `orm:"default(1)"` //角色  0：管理员，1：正常用户
}

//根据用户名和密码查询
func QueryUserByEmailAndPassword(email, psw string) (user User, err error) {

	newOrm := orm.NewOrm()
	error := newOrm.QueryTable("User").Filter("Email", email).Filter("Pwd", psw).One(&user)
	if error == orm.ErrNoRows {
		//没有查到
		err = error
		return
	}
	return
}
