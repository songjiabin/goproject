package errno

import "fmt"

//定义的错误类
//在实际开发中，一个错误类型通常包含两部分：
// Code 部分，用来唯一标识一个错误；
// Message 部分，用来展示错误信息，这部分错误信息通常供前端直接展示。
// 这两部分映射在 `errno` 包中即为 `&Errno{Code: 0, Message: "OK"}`。

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// 上面大Error下面的一个小的error
type Err struct {
	Code    int
	Message string
	Err     error  //添加了一个 err
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

//实现了Error的方法 Error();
func (err  Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
