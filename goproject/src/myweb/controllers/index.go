package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}


//展示index 
func (c *IndexController) ShowIndex() {

	c.TplName = "index.html"

}

 