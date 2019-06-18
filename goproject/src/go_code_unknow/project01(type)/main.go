package main

import "fmt"

//类型组
type (
	btye1 int32
	rune int32
	文本 string //为string 类型做个别名

)

func main() {
	var b 文本
	b = "中文类型的名字"
	fmt.Println(b)

	//赋值
	//声明并赋值一个变量
	var i int;
	i = 1
	fmt.Println(i)

	//一般用于全局变量的声明并且赋值
	var d int = 23;
	fmt.Println(d)

	//声明并赋值一个变量 之 简洁的方式
	//全局变量是不允许的 这样声明的

	c := 20
	fmt.Println(c)

	//多个局部变量的声明和赋值
	var f, g, p int = 1, 2, 3
	fmt.Println(f, g, p)


	//多个局部变量的声明和赋值2
	ab,cd,ef:=1,2,3
	fmt.Println(ab,cd,ef)

}

//全局变量的声明和赋值
var (
	a = true
	b = "123"
	c = 124
)
