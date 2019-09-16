package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"sjbproject/models"
)

type RegisterController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (this *RegisterController) ShowRegister() {
	this.TplName = "register.html"
}

//处理注册界面的逻辑
func (this *RegisterController) HandleReg() {

	this.TplName = "register.html"
	userName := this.GetString("userName")
	passWord := this.GetString("password")

	//验证账号和密码
	if userName == "" || passWord == "" {
		logs.Debug("账号或者密码不能为空")
		return
	}

	//插入数据
	newOrm := orm.NewOrm()
	user := models.User{}
	user.UserName = userName
	user.Passwd = passWord
	_, e := newOrm.Insert(&user)
	if e != nil {
		logs.Debug("插入失败", e)
		return
	}

	this.Redirect("/", 302)

}

//展示登录界面
func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

//处理登录的逻辑
func (this *LoginController) HandleLogin() {
	this.TplName = "login.html"
	userName := this.GetString("userName")
	passWord := this.GetString("password")

	if userName == "" || passWord == "" {
		logs.Debug("账号或密码不能为空")
		return
	}

	newOrm := orm.NewOrm()
	user := models.User{UserName: userName, Passwd: passWord}
	e := newOrm.Read(&user, "UserName")
	if e != nil {
		logs.Debug("读取用户信息失败", e)
		return
	}

	//判断密码
	if passWord != user.Passwd {
		logs.Debug("密码错误", e)
		return
	}

	this.Ctx.WriteString("登录成功")

}
