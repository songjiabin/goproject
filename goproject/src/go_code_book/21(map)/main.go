package main

import (
	"fmt"
)

func main() {
	a := make(map[int]int, 10)
	a[1] = 1
	a[2] = 2
	fmt.Println(a, len(a))
	func1(a)
	fmt.Println(a)
}

func func1(map2 map[int]int) {
	delete(map2, 1)
}
