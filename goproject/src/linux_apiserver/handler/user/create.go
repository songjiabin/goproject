package user

import (
	"github.com/gin-gonic/gin"
	"linux_apiserver/model"
)

//创建用户
func Create(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {

		return
	}

	model.CreateUser()

}
