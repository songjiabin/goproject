package controllers

import (
	"errors"
	"liteblog/syserror"
	"github.com/astaxie/beego/logs"
	"liteblog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type NoteController struct {
	BaseController
}

//实现了接口NestPreparer
func (this *NoteController) NextPreparse() {
	this.MustLogin()
	//如果不是管理员
	if ( this.User.Role != 0) {
		this.Abort500(errors.New("权限不足"))
	}
}

// @router /new  [get]
func (this *NoteController) Index() {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

//保存文章
// @router /save/:key  [post]
func (this *NoteController) Save() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMustString("title", "请输入标题")
	content := this.GetMustString("content", "请输入文章内容")
	logs.Info("匹配的事：->", key)
	//根据key来查询数据库中是否存在
	note, err := models.QueryNoteByKey(key)
	note.Title = title
	note.Content = content
	note.User = &this.User

	//获取摘要
	s, _ := models.GetSummary(content)
	note.Summary = s

	if err != nil && err == orm.ErrNoRows {
		//表示没有这个key 进行插入保存
		err := models.InsertNote(&note)
		if err != nil {
			this.Abort500(syserror.NewError("保存失败", err))
		}
	} else {
		//有这个key 不能插入
		this.Abort500(syserror.NewError("保存失败", err))
	}

	this.JSONOk("保存成功", "/")

}

//编辑文章
// @router /edit/:key  [get]
func (this *NoteController) Edit() {

	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByUserIdAndNoteKey(this.User.Id, key)
	if err != nil && err == orm.ErrNoRows {
		this.Abort500(syserror.NewError("文章不存在", err))
	}
	this.Data["note"] = note
	this.Data["key"] = key
	this.TplName = "note_edit.html"
}

//编辑文章后保存文章
// @router /editSave/:key  [post]
func (this *NoteController) EditSave() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetString("title")
	content := this.GetString("content")

	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.NewError("文章不存在", err))
	}

	note.Content = content
	note.Title = title

	err = models.UpdateNoteByKey(note)
	if err != nil {
		this.Abort500(syserror.NewError("更新失败", err))
	}

	this.JSONOk("保存成功", fmt.Sprintf("/details/%s", key))

}

//删除
// @router /delete/:key  [post]
func (this *NoteController) DeleteNote() {
	key := this.Ctx.Input.Param(":key")
	err := models.DeleteNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.NewError("删除失败", err))
	}

	this.JSONOk("删除成功", "/")
}
