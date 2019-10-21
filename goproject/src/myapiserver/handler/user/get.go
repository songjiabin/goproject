package user

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/model"
	"myapiserver/pkg/errno"
)

//获取具体的用户信息
func GetUser(c *gin.Context) {
	//获取用户信息
	username := c.Param("username")
	//对username进行验证
	userModel, e := model.GetUser(username)
	if e != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, userModel)

}
