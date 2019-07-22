package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
)

type RegisterController struct {
	beego.Controller
}

//点击register 进入的界面
func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

//当点击提交表单的时候
func (c *RegisterController) Post() {
	userName := c.GetString("userName")
	userPsw := c.GetString("password")

	//得到账号和密码

	//01 校验
	if userName == "" || userPsw == "" {
		beego.Info("数据不能为空")
		//重定向
		c.Redirect("/register", 302)
		return
	}

	// o2 插入数据库中去
	newOrm := orm.NewOrm()

	user := models.User{}
	user.Name = userName
	user.Pwd = userPsw

	_, e := newOrm.Insert(&user)
	if e != nil {
		c.Redirect("/register", 302)
		return
	}

	c.Ctx.WriteString("注册成功")

	beego.Info(userName, userPsw, "插入数据库成功...")
	c.TplName = "register.html"
}
