package main

import (
	"fmt"
)

func main() {

	//定义 int类型a
	a := 1
	//定义int指针类型b,保存a的内存地址
	var b *int
	b = &a
	// 打印a内存地址
	fmt.Println(b)
	//表示该内存地址所指向的内存 其实等价于 a
	*b = 555
	fmt.Println(*b, a)

	fmt.Println("------------------")

	var c *int
	c = nil

	d := 20
	c = &d
	*c = 200
	fmt.Println(d)

	//使用new
	//定义指针
	var p *int
	//这个指针指向了一块空的内存
	p = new(int)
	//给这个空内存地址上赋值为20
	*p = 20
	fmt.Println(*p)

}
