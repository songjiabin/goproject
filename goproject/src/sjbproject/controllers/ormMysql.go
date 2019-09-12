package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"sjbproject/models"
)

type OrmController struct {
	beego.Controller
}

func (c *OrmController) ShowOrm() {

}

//插入
func (c *OrmController) InsertOrm() {
	c.TplName = "orm.html"
	newOrm := orm.NewOrm()
	user := models.User{}
	user.Name = "宋佳宾"
	_, e := newOrm.Insert(&user)
	if e != nil {
		beego.Info("插入数据出错了", e)
		return
	}

}

//查询
func (c *OrmController) QueryOrm() {
	c.TplName = "orm.html"

	newOrm := orm.NewOrm()
	user := models.User{Name: "宋佳宾"}
	error := newOrm.Read(&user, "Name")
	if error != nil {
		beego.Info("查询出错了哦", error)
		return
	}

	beego.Info("查询成功", user)

}

func (c *OrmController) UpdateOrm() {
	c.TplName = "orm.html"
	newOrm := orm.NewOrm()
	user := models.User{Id: 1}
	newOrm.Read(&user, "id")
	logs.Debug(user)
	user.Name = "更改后的数据"
	_, e := newOrm.Update(&user)
	if e != nil {
		logs.Debug("更新数据失败",e)
		return
	}

	logs.Debug("更新数据成功")

}
func (c *OrmController) DeleteOrm() {
	c.TplName = "orm.html"
	newOrm := orm.NewOrm()
	user := models.User{Id: 1}
	newOrm.Read(&user, "id")
	logs.Debug(user)

	_, e := newOrm.Delete(&user)
	if e != nil {
		logs.Debug("删除数据失败",e)
		return
	}

	logs.Debug("删除数据成功")

}