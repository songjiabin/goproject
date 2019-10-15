package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"myapiserver/handler"
	"myapiserver/pkg/errno"
	"net/http"
)

func Create(c *gin.Context) {

	var r struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	if err = c.ShouldBindJSON(&r); err != nil {
		//处理
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.UserName == "" && r.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": errno.ErrBind,
		})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.UserName, r.Password)

	if r.UserName == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

}
