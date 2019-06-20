package main

import (
	"fmt"
)

func main() {

	//声明一个长度为2的string数组
	var a [2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//声明数组并赋值
	b := [4] int{1, 2, 3, 4}
	fmt.Println(b)

	d := [12]int{}
	e := d
	fmt.Println(d, e)

	//将0位置赋值1
	g := [4]int{0: 1}
	fmt.Println(g)

	//自动计算长度
	h := [...]int{3, 4, 5}
	fmt.Println(h)

	w := [...]int{24: 1}
	fmt.Println(w)

	//指针类型
	p := [...]int{24: 1}
	//p_为p的内存地址
	p_ := &p;
	fmt.Println(*p_)

	//pw 是指向数组的指针
	pw := new([10]int)
	fmt.Println(pw)

	cd := [2][3]int{{1, 2, 3}, {2, 3, 5}}
	fmt.Println(cd)

}
