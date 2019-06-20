package main

import "fmt"

//长量
const (
	//第一个 iota 表示为 0，因此 a 等于 0，
	a = iota
	//当 iota 再次在新的一行使用时，它的值增加了 1，因此 b 的值是 1。
	b = iota
	//也可以像下面这样，省略 Go 重复的 = iota：
	c
)

func main() {
	//打印结果是 0 ，1 ，2
	fmt.Println(a, b, c)
}
