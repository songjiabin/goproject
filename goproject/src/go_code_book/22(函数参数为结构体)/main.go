package main

import (
"fmt"
)

type Student struct {
	name string
	age  int
	sex  byte
}

func test(sty *Student) {
	sty.name = "改变了！"
	fmt.Println(sty)
}

func main() {

	s := &Student{name: "宋佳宾", age: 22, sex: 'c'}
	test(s)
	fmt.Println(*s)

}
