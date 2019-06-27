package main

import "fmt"

func main() {
	//空接口可以实现任意类型的数据
	//如：fmt.Println(...interface{})
	var a interface{} = 23
	fmt.Println(a)
}
