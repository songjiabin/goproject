package main

import "fmt"

//回调函数
// 函数的参数有一个参数是函数的类型，这个函数就是回调函数

type funcType func(int, int) int

//这个函数中有函数类型 这个函数就是回调函数
func test(a, b int, funcTest funcType) (result int) {
	fmt.Println("test")
	result = funcTest(a, b)
	return
}

//实现 funcType方法

func interfaceFunc(i1 int, i2 int) int {
	return i1 + i2
}

func main() {
	a := test(1, 2, interfaceFunc)
	fmt.Println(a)

	b := func(int, int) int {
		return 2
	}

	f := b(1, 1)
	fmt.Println(f)

	func() {
		fmt.Println("-------------")
	}()


	result:=func(x,y int) int{
		fmt.Println("00")
		return x-y
	}(1,2)

	fmt.Println(result)




}
