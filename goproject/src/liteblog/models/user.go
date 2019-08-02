package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int     `orm:"pk;auto" json:"Id"`
	Name   string  `orm:"unique" json:"name"	`
	Email  string  `orm:"unique" json:"email"	`
	Pwd    string  `json:"-"`
	Avatar string  `json:"avatar"`                //头像
	Role   int     `orm:"default(1)" json:"role"` //角色  0：管理员，1：正常用户
	Note   []*Note `orm:"reverse(many)"`          //一个用户下可以有多个文章
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

//根据昵称查询用户
func QueryUserByName(name string) (user User, err error) {
	newOrm := orm.NewOrm()
	error := newOrm.QueryTable("User").Filter("Name", name).One(&user)
	if error == orm.ErrNoRows {
		//没有查到
		err = error
		return
	}

	return
}

//根据邮箱查询用户
func QueryUserByEmail(email string) (user User, err error) {
	newOrm := orm.NewOrm()
	error := newOrm.QueryTable("User").Filter("Email", email).One(&user)
	if error != nil && error == orm.ErrNoRows {
		//没有查到
		err = error
		return
	} else {
		err = nil
		//查到了
	}
	return
}

// 新增用户 注册
func RegeisterUser(user *User) (err error) {
	newOrm := orm.NewOrm()
	_, err = newOrm.Insert(user)
	if err != nil {
		return
	} else {
		err = nil
		//查到了
	}
	return nil
}


