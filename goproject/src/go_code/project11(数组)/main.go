package main

import (
	"fmt"
)

func main() {

	//声明一个长度为2的string数组
	var a [2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//声明数组并赋值
	b := [4] int{1, 2, 3, 4}
	fmt.Println(b)



}
