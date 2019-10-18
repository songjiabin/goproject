package user

import "myapiserver/model"

type CreateRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type CreateResponse struct {
	Username string `json:"username" form:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	offset   int    `json:"offset"`   //从哪里开始获取
	Limit    int    `json:"limit"`    //获取多少个
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
