package errno

type Errno struct {
	Code    int
	Message string
}

//实现了错误的接口
func (err *Errno) Error() string {
	return "实现了Error接口"
}

//解析
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Message
	default:
		return UnknownError.Code, typed.Error()
	}
}
