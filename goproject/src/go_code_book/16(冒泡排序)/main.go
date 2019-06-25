package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := [10]int{}
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}

	//从小到大
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				num := a[i]
				a[i] = a[j]
				a[j] = num
			}
		}
	}

	fmt.Println(a)

	//从小到大
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				num := a[i]
				a[i] = a[j]
				a[j] = num
			}
		}
	}

	fmt.Println(a)
}
