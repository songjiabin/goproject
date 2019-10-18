package user

import (
	"github.com/gin-gonic/gin"
	"myapiserver/handler"
	"myapiserver/pkg/errno"
	"myapiserver/service"
)

//获取所有的user用户列表
func List(c *gin.Context) {
	//http://localhost:port/user?username=xxx&offset=x&limit=x
	var listRequest ListRequest
	if err := c.Bind(&listRequest); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.Listuser(listRequest.Username, listRequest.offset, listRequest.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: uint64(count),
		UserList:   infos,
	})
}
