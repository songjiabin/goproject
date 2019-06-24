package test

import "fmt"

func init() {
	fmt.Println("init..>test")
}

func Test(a, b int) int {
	return a + b
}
