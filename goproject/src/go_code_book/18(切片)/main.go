package main

import (
	"fmt"
)

func main() {

	makeClice()

	a := [...]int{1, 2, 3, 4, 5, 6}
	// a[start,end,max]
	// start:起始位置 1
	// end: 结束位置 2
	// max: 容量 = 6-1
	// len: end-start 3-1
	slipce := a[1:3:6]
	fmt.Println(slipce, len(slipce), cap(slipce))

	fmt.Println("------------------")

	b := []int{}
	b = append(b, 1)
	fmt.Println(b, len(b), cap(b))

	rapChange()
}

//制作切片
func makeClice() {

	//制作一个切片  长度是5 容量是10（mei）
	a := make([]int, 5, 10)
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := b[:]

	fmt.Println(c)

	fmt.Println("****************************")
}

//切片容量的变化
func rapChange() {
	fmt.Println("****************************")


	a := make([]int, 0, 1)
	olderCap := cap(a)
	fmt.Println("容量cap:", olderCap)
	for i := 0; i < 8; i++ {
		a = append(a, i)
		if newCap := cap(a); newCap > olderCap {
			fmt.Println("容量cap:", newCap)
			olderCap = newCap
		}

	}

}
