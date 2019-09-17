package utils

import "github.com/astaxie/beego/logs"

func HandleError(msg string, err error, ) {
	if err != nil {
		logs.Debug(msg, err)
		return
	}
}
