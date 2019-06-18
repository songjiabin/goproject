package main

import (
	"fmt"
)

//一个结构体（struct）就是一组字段（field）。

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{x: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	v4 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	//打印结构体
	fmt.Println(Vertex{1, 2,})
	// 访问结构体的数据
	v := Vertex{1, 2,}
	fmt.Println(v.x)

	//结构体和指针
	p := &v             //p等于v结构体的内存地址
	fmt.Println((*p).x) //可以通过*p（也就是v来操作v结构体）
	//当然也可以写成 =>
	fmt.Println(p.x) // (*p).x 这样写太啰嗦了，所有使用隐士的简介引用

	fmt.Println(v1, v2, v3, v4)

}
