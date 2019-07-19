package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type PostController struct {
	beego.Controller
}

type OrmController struct {
	beego.Controller
}

type OrmSelectController struct {
	beego.Controller
}

type UpdateController struct {
	beego.Controller
}

type DeleteController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *TestController) Get() {
	c.Data["data"] = "Learn beego now"
	c.TplName = "test.html"
}

// post请求
func (c *PostController) Post() {
	c.Data["data"] = "Learn beego post now"
	c.TplName = "test.html"
}

// orm 插入操作
func (c *OrmController) Get() {
	//1.有ORM对象
	orm := orm.NewOrm()
	//2.要插入一个数据的结构体对象 并对结构体赋值
	user := models.User{
		Name: "基本密码",
		Pwd:  "123",
	}
	//3.插入
	_, e := orm.Insert(&user)
	if e != nil {
		fmt.Println("插入失败", e)
		return
	}
	fmt.Println("插入成功")

	c.Data["data"] = "Learn orm insert content"
	c.TplName = "test.html"
}

//查询数据
func (c *OrmSelectController) Get() {

	// 创建orm对象
	newOrm := orm.NewOrm()

	// 01 根据指定的id查询数据
	//user := selectById(newOrm)

	// 02 根据指定的列名来查
	user := selectByName(newOrm)

	c.Data["name"] = user.Name
	c.TplName = "user.html"
}

//根据指定的id进行查询
func selectById(newOrm orm.Ormer) (user models.User) {
	// 查询的对象
	user = models.User{}

	// 指定查询的字段
	user.Id = 1 //查询id为1的数据

	// 查询
	err := newOrm.Read(&user)
	if err != nil {
		beego.Info("查询失败")
		return
	}
	beego.Info("查询成功：", user)

	return
}

//根据指定的名字来查询
func selectByName(newOrm orm.Ormer) (user models.User) {
	user = models.User{}

	// 指定查询的字段
	user.Name = "基本密码" //查询id为1的数据

	// 查询 根据名字来查
	err := newOrm.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败")
		return
	}
	beego.Info("查询成功：", user)

	return
}

//更新user
func (c *UpdateController) Get() {
	//1
	newOrm := orm.NewOrm()
	//2
	user := models.User{}

	//3
	user.Id = 1

	//4
	e := newOrm.Read(&user)
	if e != nil {
		fmt.Println(e)
		return
	} else {
		user.Name = "基本密码(更新后)"
		user.Pwd = "123456"
	}

	//5
	newOrm.Update(&user)

	c.Data["name"] = user.Name
	c.TplName = "user.html"
}

//删除
func (c *DeleteController) Get() {

	newOrm := orm.NewOrm()
	user := models.User{}
	user.Id = 2
	_, e := newOrm.Delete(&user)
	if e != nil {
		beego.Info("删除成功")
		return
	}

	c.Data["name"] = "删除成功"
	c.TplName = "user.html"
}

func (c *MainController) showLogin() {
	c.TplName = "login.html"
}
