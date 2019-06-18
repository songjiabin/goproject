package main

import (
	"fmt"
	_ "math"
)

/**
  *	Go 没有类。不过你可以为结构体类型定义方法。
 */
type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	//调用的时候 可以用接受者直接调用方法
	fmt.Println(v.Abs(v))
}

// Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
// 方法只是一个带有接受者参数的函数
func (v Vertex) Abs(a Vertex) float64 {
	return v.X + v.Y + a.X;
}
