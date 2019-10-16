package err

type MyError struct {
	Code    int
	Message string
}

//实现了Erro接口
func (MyError *MyError) Error() string {
	return ""
}

type MyErrorOne struct {
	Code    int
	Message string
	MyError error
}

//实现了Erro接口
func (myErrorOne *MyErrorOne) Error() string {
	return ""
}

func New(myError *MyError, err error) error {
	return &MyErrorOne{
		Code:    myError.Code,
		Message: myError.Message,
		MyError: err,
	}
}

func (err *MyErrorOne) Add(messsage string) error {
	err.Message += " " + messsage
	return err
}
