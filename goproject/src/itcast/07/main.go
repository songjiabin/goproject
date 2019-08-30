package main

import "fmt"

func main() {

	olderInts := []int{
		1, 2, 3, 4, 5,
	}
	//必须知道newInts的长度
	newInts := make([]int, len(olderInts))
	copy(newInts, olderInts)

	fmt.Println(newInts)

	test(olderInts)
	fmt.Println(olderInts)
}

func test(arr []int) {

	arr[0] = 10
	fmt.Println(arr)

}
