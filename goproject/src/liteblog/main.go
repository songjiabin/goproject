package main

import (
	_"liteblog/models"
	_ "liteblog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initTemplate()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", equrl)
}

//比较两个url
func equrl(url1, url2 string) (result bool) {
	url1_1 := strings.Trim(url1, "/")
	url2_2 := strings.Trim(url2, "/")
	return strings.Compare(url1_1, url2_2) == 0
}
