package router

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler/sd"
	"myapiserver/handler/user"
	"net/http"
	"time"
)

//中间项
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	// 添加请求头
	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，
	// 需要通过 `gin.Recovery()`来恢复 API 服务器
	g.Use(gin.Recovery())
	//强制浏览器不使用缓存
	g.Use(NoCache)
	//浏览器跨域 OPTIONS 请求设置
	g.Use(Options)
	//一些安全设置
	g.Use(Secure)
	//自己设置的中间项
	g.Use(mw...)

	//404的处理
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route")
	})

	//对请求的api进行分组
	group := g.Group("sd")

	{
		group.GET("/health", sd.HealthCheck)
		//group.GET("/disk", sd.DiskCheck)
		//group.GET("/cpu", sd.CPUCheck)
		//group.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create_relase)     //创建用户
		//u.POST("/:id", user.Delete) //删除用户
		//u.POST("/:id", user.Update) //更新用户
		//u.GET("", user.List)        //获取用户列表
		//u.GET("/:username",user.Get)         //获取指定用户的详细信息
	}

	return g
}

//以防止客户端缓存HTTP响应。
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

//浏览器跨域 OPTIONS 请求设置
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//一些安全设置
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
