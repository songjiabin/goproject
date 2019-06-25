package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

//数组指针做函数的参数
//当直接传递过去的时候 原始的值是不变的
func modify(array [5]int) {
	array[0] = 100
}

//可以传递地址
func modify2(array *[5]int) {

	(*array)[0] = 100

}
