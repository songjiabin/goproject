package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//测试连接
func Test(c *gin.Context) {
	c.String(http.StatusOK, "连接成功")
}

