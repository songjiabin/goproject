package user

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/model"
	"myapiserver/pkg/auth"
	"myapiserver/pkg/errno"
	"myapiserver/pkg/token"
)

//登录
func Login(c *gin.Context) {

	var u model.UserModel
	//绑定实体
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//根据姓名获取实体
	userModel, e := model.GetUser(u.Username)
	if e != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	//比较密码
	if err := auth.Compare(userModel.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	//生成秘钥
	t, err := token.Sign(c, token.Context{ID: userModel.Id, Username: userModel.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{
		Token: t,
	})

}
