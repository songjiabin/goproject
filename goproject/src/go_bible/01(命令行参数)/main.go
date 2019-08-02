package main

import (
	"os"
	"fmt"
	"strings"
)

//1、go build main.go   前提下，定位到该目录下
//2、对生成的main.exe 操作  main.ext 1 3 4 5 6
//3、打印的结果是： 1 3 4 5 6

func main() {
	var s, sep string
	//第一个参数是命令  os.Args的第一个元素，os.Args[0], 是命令本身的名字
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	a := 1
	a++

	//这样写事非法的,因为a++是语句，而不是表达式
	//b=a++
	fmt.Println(a)

	//无限循环
	//for true{
	//	fmt.Println("....")
	//}

	//声明变量的方式
	// 短变量声明，最简洁，但只能用在函数内部，而不能用于包变量
	//bc := ""
	////依赖于字符串的默认初始化零值机制，被初始化为""
	//var bc string
	////
	//var bc = ""
	//var bc string = ""
	//
	//fmt.Println(bc)

	array := []string{
		"1",
		"2",
		"3",
	}
	newArrayString:= strings.Join(array, "----")
	fmt.Println(newArrayString)












}
