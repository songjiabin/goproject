package controllers

type ErrorController struct {
	BaseController
}

//处理404的错误
func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found---404"
	this.TplName = "error/404.html"
}

//处理500的错误
func (this *ErrorController) Error500() {
	err, _ := this.Data["error"].(error)
	this.Data["err"] = err
	this.TplName = "error/500.html"
}
