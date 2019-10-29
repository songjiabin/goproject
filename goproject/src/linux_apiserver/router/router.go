package router

import (
	"github.com/gin-gonic/gin"
	"linux_apiserver/handler"
	"linux_apiserver/handler/user"
)

func Load(g *gin.Engine) *gin.Engine {

	//设置拦截器
	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，
	// 需要通过 `gin.Recovery()`来恢复 API 服务器
	g.Use(gin.Recovery())

	//测试连接
	g.GET("/test", handler.Test)
	group := g.Group("/v1/user")
	{
		group.POST("", user.Create)
		group.POST("",user.Delete)
	}
	return g
}
