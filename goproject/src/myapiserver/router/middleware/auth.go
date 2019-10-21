package middleware

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/pkg/errno"
	"myapiserver/pkg/token"
)

func AuthMiddleware(c *gin.Context) {

	// Parse the json web token.
	if _, err := token.ParseRequest(c); err != nil {
		logs.Info(err)
		handler.SendResponse(c, errno.ErrTokenInvalid, nil)
		c.Abort()
		return
	}
	c.Next()
}
