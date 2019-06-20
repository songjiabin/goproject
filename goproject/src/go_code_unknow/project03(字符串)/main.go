package main

import (
	"fmt"
	"strconv"
)

//字符串在 Go 中是 UTF-8 的由双引号（”）包裹的字符序列。如果你使用单引号（’）则
//表示一个字符（UTF-8 编码）——这种在 Go 中不是 string
func main() {

	//一旦给变量赋值，字符串就不能修改了：在 Go 中字符串是不可变的。从 C 来的用户，
	//下面的情况在 Go 中是非法的
	var a = "string"

	//得到a中的第一个元素
	b := a[1];

	//类型转换 将 a=>[]run
	c := []rune(a)

	//修改数组的第一个元素
	c[1] = 'c'

	fmt.Println(string(b), string(c))

	//字符串的拼接
	s := "abc" +
		"cdf"

	fmt.Println(s)

	//拼接负责的字符串

	str := `

    var a = "string"

	//得到a中的第一个元素
	b := a[1];

	//类型转换 将 a=>[]run
	c := []rune(a)

	//修改数组的第一个元素
	c[1] = 'c'

	fmt.Println(string(b), string(c))

	//字符串的拼接 
	s := "abc" +
		"cdf"

	fmt.Println(s)

	`

	fmt.Println(str)

	a65 := 65;
	fmt.Println(string(a65))
	fmt.Println(strconv.Itoa(a65))
	bcd := strconv.Itoa(a65)
	fmt.Println(strconv.Atoi(bcd))

}
