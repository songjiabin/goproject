package controllers

import (
	"DDFruit/models"
	"regexp"
	"strings"
)

type UserControllers struct {
	BaseController
}

//  @router /register  [get]
//展示注册的界面
func (this *UserControllers) ShowRegister() {
	this.TplName = "register.html"
}

//  @router /register  [post]
//处理注册的界面
func (this *UserControllers) HandleRegister() {
	this.TplName = "register.html"
	user_name := this.GetString("user_name")
	user_pwd := this.GetString("pwd")
	user_cpwd := this.GetString("cpwd")
	user_email := this.GetString("email")

	if user_name == "" || user_pwd == "" || user_cpwd == "" || user_email == "" {
		this.Data["errmsg"] = "数据不完整，请重新注册"
		return
	}

	if strings.Compare(user_pwd, user_cpwd) != 0 {
		this.Data["errmsg"] = "两次输入密码不一致，请重新注册"
		return
	}

	//校验邮箱
	reg, _ := regexp.Compile("^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	res := reg.FindString(user_email)
	if res == "" {
		this.Data["errmsg"] = "邮箱格式不正确"
		return
	}

	//插入数据
	user := &models.User{
		UserName: user_name,
		PassWord: user_pwd,
		Email:    user_email,
	}
	err := user.InsertUser()
	if err != nil {
		return
	}
}

// @router /login  [get]
//展示登录的界面
func (this *UserControllers) ShowLogin() {
	this.TplName = "login.html"
}

// @router /login [post]
//处理登录的逻辑
func (this *UserControllers) HandleLogin() {
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	if username == "" || pwd == "" {
		this.Data["errmsg"] = "登录数据不完整，请重新输入！"
		this.TplName = "login.html"
		return
	}



}
