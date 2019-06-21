package main

import (
	"fmt"
)

// fun ->匿名函数
func main() {
	a := func() {
		fmt.Println("Fucn A")
	}

	a()

	//调用闭包
	f := closre(1)
	fmt.Println(f(1))

}

func closre(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
