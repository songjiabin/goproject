package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

//登录
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

//登录界面
func (c *LoginController) ShowLogin() {
	c.TplName = "login.html"
}

//出来post登录接口
func (c *LoginController) HandleLogin() {


	c.TplName = "login.html"
	
	userName := c.GetString("userName")
	userPsw := c.GetString("userPsw")
	//非空判断
	if userName == "" || userPsw == "" {
		c.Redirect("/login", 302)
		return
	}

	//查数据中
	if isHaveData := findUser(userName, userPsw); !isHaveData {
		beego.Info("查不到此数据")
		return 
	} 

	beego.Info(userName, userPsw)
	c.Ctx.WriteString("登录成功")
	
}

func findUser(userName, userPsw string) (isHave bool) {
	isHave = false

	newOrm := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = userPsw

	e := newOrm.Read(&user, "Name", "Pwd")
	if e != nil {
		beego.Info("错误是：", e)
		return
	} else {
		isHave = true
	}

	beego.Info("查询到的数据是：", user)
	return
}
