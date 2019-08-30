package main

import "fmt"

func test(a int) {
	a++
	fmt.Println(a)
}

func mainaaa() {

	a := 0
	for i := 0; i < 10; i++ {
		test(a)
	}

	fmt.Println("------------------")

	f := test2()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func test2() (func() int) {
	var a int
	return func() int {
		a++
		return a
	}
}

//阶乘
var sum int = 1

func test11111(num int) {

	if num == 1 {
		return
	}

	test11111(num - 1)
	sum *= num

}

func main() {
	test11111(5)
	fmt.Println(sum)
}
