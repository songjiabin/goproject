package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
)

type OneToOneController struct {
	beego.Controller
}

func (this *OneToOneController) Get() {

	//o := orm.NewOrm()
	//
	//p := models.PersonName{Title: "one2one title"}
	//u := models.Person{Name: "Person"}
	//
	//p.Person = &u
	//u.PersonName = &p
	//
	//o.Insert(&p)
	//o.Insert(&u)

	orm.Debug = true
	newOrm := orm.NewOrm()
	var person []models.Person

	newOrm.QueryTable("Person").Filter("PersonName__Title","myTitle").All(&person)

	for _, v := range person {
		this.Ctx.WriteString(v.Name+"<br/>")
	}

	this.TplName = "one2one.html"

}
