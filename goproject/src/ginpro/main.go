package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//https://www.yoytang.com/go-gin-doc.html#Gin%20的介绍
func main() {
	// 初始化引擎
	engine := gin.Default()
	//注册一个路由和处理函数
	engine.Any("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello_world_songjiabin")
	})

	engine.GET("/name", func(context *gin.Context) {
		context.String(http.StatusOK, "get请求哦")
	})

	engine.Any("/any/good", hanlderGet)
	engine.Run(":9217")

}

func hanlderGet(context *gin.Context) {
	context.String(http.StatusOK, "请求成功200")
}
