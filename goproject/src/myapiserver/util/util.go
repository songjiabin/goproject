package util

import (
	"github.com/gin-gonic/gin"
	"github.com/ventu-io/go-shortid"
)

//获取请求的X_request-id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

//唯一非顺序短ID的生成器
func GenShortId() (string, error) {
	return shortid.Generate()
}
