package controllers

type IndexController struct {
	//继承自它
	BaseController
}

// @router /  [get]
func (this *IndexController) Index() {
	this.TplName = "index.html"
}

// @router /about  [get]
func (this *IndexController) IndexAbout() {
	this.TplName = "about.html"
}

// @router /message  [get]
func (this *IndexController) IndexMessage() {
	this.TplName = "message.html"
}
