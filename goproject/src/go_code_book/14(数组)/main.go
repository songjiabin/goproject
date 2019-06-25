package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}

	fmt.Println(a)

	//的有多少的[]就是多少维数组

	array := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 1, 2, 3},
	}

	fmt.Println(array)




}
