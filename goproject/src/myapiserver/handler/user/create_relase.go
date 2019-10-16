package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"myapiserver/handler"
	"myapiserver/model"
	"myapiserver/pkg/errno"
	"myapiserver/util"
)

// 创建一个新的连接  用户请求参数
func Create_relase(c *gin.Context) {
	log.Info("User create funtion called .", lager.Data{"X-Rquest-ID": util.GetReqID(c)})

	//绑定请求
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//绑定数据和实体（orm的方式）
	userModel := model.UserModel{
		Username: r.Username,
		Password: r.Password,
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

	//插入数据到数据库中
	error := userModel.Create()
	if error != nil {
		log.Info("数据库错误为:--->" + error.Error())
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	//插入成功了
	rsp := CreateResponse{
		Username: r.Username,
	}
	//ok
	handler.SendResponse(c, nil, rsp)

}
