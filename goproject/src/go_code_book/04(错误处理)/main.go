package main

import "fmt"

func main() {
	test()
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("...", err)
		}
	}()
	divide(5, 0)
	fmt.Println("end of test")
}

func divide(a, b int) int {
	return a / b
}
