package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"linux_apiserver/errno"
	"linux_apiserver/handler"
	"linux_apiserver/model"
)


//创建用户
func Create(c *gin.Context) {
	fmt.Println("come on ")
	var user model.User
	if err := c.Bind(&user); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	 //验证数据
	if err := user.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//将用户的密码加密
	if err := user.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	log.Infof("user名字 %s",user.Username)
	log.Infof("user密码 %s",user.Password)
	
	if err := model.CreateUser(&user); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	//创建成功
	handler.SendResponse(c, nil, gin.H{
		"username": user.Username,
	})
}
