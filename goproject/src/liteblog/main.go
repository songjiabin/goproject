package main

import (
	_ "liteblog/models"
	_ "liteblog/routers"
	"github.com/astaxie/beego"
	"strings"
	"encoding/gob"
	"liteblog/models"
)

func main() {
	initSession()
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

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	//设置 Session 的引擎，默认是 memory，目前支持还有 file、mysql、redis 等，配置文件对应的参数名：sessionprovider。
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//设置对应 file、mysql、redis 引擎的保存路径或者链接地址，默认值是空，配置文件对应的参数：sessionproviderconfig。
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
