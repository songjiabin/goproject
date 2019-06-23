package main

import (
	"fmt"
)

// 定义了一个函数类型  将其命名为MyFuncType
// 接受一个int类型的参数，并返回一个bool类型的结果
type MyFuncType func(int) bool

func isBigThan5(n int) bool {
	return n > 5
}

func display(arr []int, f MyFuncType) {
	for _, v := range arr {
		if f(v) {
			fmt.Println(v)
		}
	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	compartBigOrSmal(arr, isBigThan5)
}

func compartBigOrSmal(array []int, f MyFuncType) {
	for _, v := range array {
		if f(v) {
			fmt.Println(v)
		}
	}
}
