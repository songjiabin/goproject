package user

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/model"
	"myapiserver/pkg/errno"
	"strconv"
)

//删除用户
func Delete(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.UnknownError, nil)
		return
	}

	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
