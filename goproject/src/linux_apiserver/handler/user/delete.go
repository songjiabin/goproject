package user

import (
	"github.com/gin-gonic/gin"
	"linux_apiserver/handler"
	"linux_apiserver/model"
	"strconv"
)

//删除用户信息
func Delete(c *gin.Context) {
	deleteId := c.PostForm("id")
	id, err := strconv.Atoi(deleteId)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	//删除用户
	if err := model.DeleteUser(id); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	//删除成功
	handler.SendResponse(c, nil, nil)

}
