package main

import (
	"errors"
	"fmt"
)

func main() {
	fun1()

	result, error := myDiv(6, 5)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("结果是：", result)
	}

}

//平常的使用
func fun1() {
	err := errors.New("错误了")
	fmt.Println(err)
}

//两个数相除
func myDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}
