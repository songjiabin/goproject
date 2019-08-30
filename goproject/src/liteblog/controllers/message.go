package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
)

// 评论处理的控制器
type MessageController struct {
	BaseController
}

//@router  /new/:key  [post]
func (this *MessageController) NewMessage() {
	// 新增评论，必须用户登陆
	this.MustLogin()
	key := this.Ctx.Input.Param(":key")
	// 获取评论的内容
	content := this.GetMustString("content", "请输入内容！")
	// 新增一个评论的唯一key
	k := this.UUID()

	//如果数据中存在该key那么就不让保存
	_, err := models.QueryMessageByKey(k)
	if err == nil {
		this.Abort500(syserror.NewError("不能重复添加", err))
	}

	//根据key查询文章
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.NewError("保存失败", err))
	}

	message := models.Message{
		Content: content,
		Note:    &note,
		Key:     k,
		User:    &this.User,
	}
	errMessage := models.SaveMessage(message)
	if errMessage != nil {
		this.Abort500(syserror.NewError("保存失败", err))
	}
	// 提示保存成功
	this.JSONOk("保存成功！", "/");
}

//留言功能
//@router  /lnew  [post]
func (this *MessageController) NewLMessage() {
	// 新增评论，必须用户登陆
	this.MustLogin()
	content := this.GetMustString("content", "请输入内容！")
	// 新增一个评论的唯一key
	k := this.UUID()
	_, err := models.QueryLMessageByKey(k)
	if err == nil {
		this.Abort500(syserror.NewError("不能重复添加", err))
	}

	message := models.LMessage{
		Content: content,
		Key:     k,
		User:    &this.User,
	}
	errMessage := models.SaveLMessage(message)
	if errMessage != nil {
		this.Abort500(syserror.NewError("保存失败", err))
	}
	// 提示保存成功
	this.JSONOk("保存成功！", "/message");
}
