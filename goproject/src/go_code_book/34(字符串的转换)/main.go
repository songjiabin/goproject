package main

import (
	"fmt"
	"strconv"
)

//字符串的转换

func main() {

	//转换成字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//1234 要追加的数  10：10进制
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcdefghijk") //追加字符穿
	fmt.Println(string(slice))

	// 其他类型转换成字符串
	var str string
	str = strconv.FormatBool(true)
	fmt.Println(str)
	// 'f'：打印格式（小数方式） -1：小数点后面位数 64以float64
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)

	//整形转字符串
	bd := strconv.Itoa(34)
	fmt.Println(bd)

	//字符串正常整形
	my := "fd"
	newMy, error := strconv.Atoi(my)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%T\n", newMy)
	}

	//字符串转其他类型
	bol, error := strconv.ParseBool("true")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(bol)
	}

}
