package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//全局请求的中间件
//所有的请求和返回值都进行统一的处理
// 在请求和返回的 Header 中插入 `X-Request-Id`（`X-Request-Id` 值为 32 位的 UUID，用于唯一标识一次 HTTP 请求）
// 日志记录每一个收到的请求
func RequestId(c *gin.Context) {
	requestId := c.Request.Header.Get("X-Request-Id")
	// Create request id with UUID4
	if requestId == "" {
		//X-Request-Id
		u4, _ := uuid.NewV4()
		requestId = u4.String()
	}
	c.Writer.Header().Set("X-Request-Id", requestId)
	c.Next()
}
