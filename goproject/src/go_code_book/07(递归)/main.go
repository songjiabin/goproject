package main

import (
	"fmt"
)

func test() (sum int) {

	sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return

}

func test2(i int) int {
	//使用递归
	if i == 1 {
		return 1
	}

	return i + test2(i-1)

}

func main() {
	fmt.Println(test())
	fmt.Println(test2(100))
}
