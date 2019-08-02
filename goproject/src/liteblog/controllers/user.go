package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
	"strings"
	"errors"
)

type UserController struct {
	//继承自它
	BaseController
}

// @router /login  [post]
func (this *UserController) Login() {

	//获取数据
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")

	var (
		user models.User
		err  error
	)
	//根据邮箱和密码查询数据
	if user, err = models.QueryUserByEmailAndPassword(email, password); err != nil {
		// About500使在BaseController里面定义的，使对beego的ctx.Abort(..)的封装，目的：优化错误提示。
		this.Abort500(syserror.NewError("邮箱或密码不对", err))
	}

	//将用户信息保存到session上
	this.SetSession(SESSION_USER_KEY, user)
	//JSONOk使在BaseController里面定义的，使对beego的ctx.ServeJSON()的封装，目的：优化提示
	//这里将返回 {code: 0, msg: "登陆成功", action: "/"}
	this.JSONOk("登陆成功", "/")
}

// @router /register  [post]
func (this *UserController) Register() {
	name := this.GetMustString("name", "昵称不能为空！")
	email := this.GetMustString("email", "邮箱不能为空！")
	pwd1 := this.GetMustString("password", "密码不能为空！")
	pwd2 := this.GetMustString("password2", "确认密码不能为空！")
	if strings.Compare(pwd1, pwd2) != 0 {
		this.Abort500(errors.New("密码与确认密码 必须要一致！"))
	}

	if _, err := models.QueryUserByName(name); err == nil {
		this.Abort500(syserror.NewError("用户昵称已经存在!", err))
	}

	if _, err := models.QueryUserByEmail(email); err == nil {
		this.Abort500(syserror.NewError("用户邮箱已经存在！", err))
	}

	user := models.User{
		Name:  name,
		Email: email,
		Pwd:   pwd1,
		Role:  1,//默认为1
	}

	//如果通过了 说注册成功 进行保存
	err := models.RegeisterUser(&user)
	if err != nil {
		this.Abort500(syserror.NewError("用户注册失败", err))
	}

	this.JSONOk("注册成功", "/login")

}

// @router /logout [get]
func (this *UserController) Logout() {
	//退出登录
	this.DelSession(SESSION_USER_KEY)
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
	this.Redirect("/", 302)
}
