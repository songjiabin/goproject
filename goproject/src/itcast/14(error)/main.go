package main

import (
	"errors"
	"fmt"
)

func test(a, b int) (value int, err error) {

	if b == 0 {
		err = errors.New("除数不能为0")
		return
	} else {
		value = a / b
		err = nil
		return
	}

}

func main() {

	value, err := test(2, 0)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(value)
	}

}
