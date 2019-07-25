package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
	"github.com/astaxie/beego/orm"
	"time"
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

	userName := c.Ctx.GetCookie("userName")
	//说明找到了userName
	c.Data["userName"] = userName

	c.TplName = "login.html"
}

//出来post登录接口
func (c *LoginController) HandleLogin() {

	c.TplName = "login.html"

	userName := c.GetString("userName")
	userPsw := c.GetString("password")
	//非空判断
	if userName == "" || userPsw == "" {
		c.Redirect("/login", 302)
		return
	}

	//查数据中
	if isHaveData := findUser(userName, userPsw); !isHaveData {
		return
	}

	//是否记住密码
	isRmberName := c.GetString("remember")
	if isRmberName == "on" {
		//记住用户名字了要
		//设置cookie 时间为1小时
		c.Ctx.SetCookie("userName", userName, time.Second*1000, "/");
	} else {
		//否则消除cookie
		c.Ctx.SetCookie("userName", "", -1, "/");
	}

	//存入session
	c.SetSession("userName", userName)

	//c.Ctx.WriteString("登录成功")

	//当登录成功后要进入index.html的界面
	c.Redirect("/index", 302)
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
