package main

import (
	_ "myweb/models"
	_ "myweb/routers"
	"github.com/astaxie/beego"
)

func main() {
	//注册 进行和index.html相对应
	beego.AddFuncMap("showPrePage", HandlePrePage)
	beego.AddFuncMap("showNextPage", HanleNextPage)
	beego.Run()
}

//当index.html中点击上一页的时候
func HandlePrePage(data int) (pageIndex int) {
	pageIndex = data - 1
	if pageIndex <= 0 {
		pageIndex = 1
	}
	return
}

//当index.html中点击下一页的时候
func HanleNextPage(data int) (pageIndex int) {
	pageIndex = data + 1
	return
}
