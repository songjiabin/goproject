package user

import (
	"github.com/gin-gonic/gin"
	"linux_apiserver/errno"
	"linux_apiserver/handler"
	"strconv"
)

//删除用户信息
func Delete(c *gin.Context) {

	deleteId := c.Param("id")
	id, err := strconv.Atoi(deleteId)
	if err != nil {
		handler.SendResponse(c,errno.UnknownError,nil)
		return
	}




}
