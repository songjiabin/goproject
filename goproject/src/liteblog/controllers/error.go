package controllers

import (
	"liteblog/syserror"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

//处理404的错误
func (c *ErrorController) Error404() {

	if c.IsAjax() {
		c.jsonerror(syserror.Error404{})
	}

	c.TplName = "error/404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		//类型判断失败后
		err = syserror.NewError("未知错误", nil)
	}

	//如果是系统定义的错误
	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.NewError(err.Error(), nil)
	}

	if (serr.ReasonError() != nil) {
		logs.Info(serr.Error(), serr.ReasonError())
	}

	if c.IsAjax() {
		c.jsonerror(serr)
	} else {
		c.Data["content"] = serr.Error()

	}
}

func (c *ErrorController) jsonerror(serr syserror.Error) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code": serr.Code(),
		"msg":  serr.Error(),
	}
	c.ServeJSON()
}
