package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	//继承自它
	BaseController
}

// @router /login  [post]
func (this *UserController) Login() {

	logs.Info("出发我了。。。。")

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
	this.SetSession(SESSION_USER_KEY, &user)
	//JSONOk使在BaseController里面定义的，使对beego的ctx.ServeJSON()的封装，目的：优化提示
	//这里将返回 {code: 0, msg: "登陆成功", action: "/"}
	this.JSONOk("登陆成功", "/")
}
