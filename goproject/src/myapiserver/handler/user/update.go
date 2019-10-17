package user

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/model"
	"myapiserver/pkg/errno"
	"strconv"
)

//更新数据
func Update(c *gin.Context) {
	//http:192.168.2.133:1259/v1/user/2
	//获取后面的2
	userId, _ := strconv.Atoi(c.Param("id"))
	userModel := model.UserModel{}
	if err := c.Bind(&userModel); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//验证数据
	if err := userModel.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//加密对密码
	if error := userModel.Encrypt(); error != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	userModel.Id = uint(userId)
	if err := model.UpdateUser(&userModel); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
