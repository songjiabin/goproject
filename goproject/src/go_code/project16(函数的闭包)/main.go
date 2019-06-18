package main

import (
	"fmt"
)

func main() {
	pos, nes := adder(), adder()


	//fmt.Println(pos(1))
	//fmt.Println(nes(2))

	for i := 0; i < 100; i++ {
		fmt.Println(pos(i), nes(i))
	}

}

//闭包
// 闭包是一个函数值
// 它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
// 函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum;
	}
}
