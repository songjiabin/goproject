package main

import (
	"fmt"
	_ "math"
)

//定义 结构体
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y

}

//接受者 指针类型的 结构体
func (v *Vertex) Scale() {

}

func main() {

	//a := Vertex{
	//	X: 1,
	//	Y: 1,
	//}
	//fmt.Println(a.Abs())

	fmt.Println(a, b, c)

}

const (
	//第一个 iota 表示为 0，因此 a 等于 0，
	a = iota
	//当 iota 再次在新的一行使用时，它的值增加了 1，因此 b 的值是 1。
	b =iota
	//也可以像下面这样，省略 Go 重复的 = iota：
	c
)
