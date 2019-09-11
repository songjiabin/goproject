package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type IndexController struct {
	beego.Controller
}

type ApiController struct {
	beego.Controller
}

type NumController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "main---get"
	c.TplName = "1.html"
}

func (c *IndexController) Post() {
	c.Data["index"] = "index"
	c.Data["Website"] = "index--->post"
	c.TplName = "1.html"
}

func (c *IndexController) GetMethod() {
	c.Data["index"] = "getMethod"
	c.Data["Website"] = "getMethod_website"
	c.TplName = "1.html"
}

func (c *IndexController) PostMethod() {
	c.Data["index"] = "getMethod"
	c.Data["Website"] = "getMethod_website"
	c.TplName = "1.html"
}

func (c *ApiController) GetApi() {
	s := c.GetString(":id")
	beego.Info(s)
	c.TplName = "1.html"
}

func (c *NumController) GetNum() {
	s := c.GetString(":id")
	beego.Info(s)
	c.TplName = "1.html"
}
