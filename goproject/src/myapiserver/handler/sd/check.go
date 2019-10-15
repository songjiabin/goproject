package sd

//业务处理的代码

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}
