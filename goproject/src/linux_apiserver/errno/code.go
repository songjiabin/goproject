package errno

var (
	OK                  = &Errno{Code: 0, Message: "成功",}
	InternalServerError = &Errno{Code: 10001, Message: "网络错误"}
	ErrBind             = &Errno{Code: 10002, Message: "绑定数据错误"}
	UnknownError        = &Errno{Code: 10003, Message: "未知错误"}

	//验证失败类型的错误
	ErrValidation = &Errno{Code: 20001, Message: "验证失败"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "签署JSON网络令牌时发生错误"}

	// 用户类型的 errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "userName没有查询到"}
	ErrEncrypt           = &Errno{Code: 20101, Message: "加密用户密码时发生错误"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "Token失效"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误"}
)
