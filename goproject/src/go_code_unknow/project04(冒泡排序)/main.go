package main

import "fmt"

//实现冒泡排序
func main() {
	a := [...]int{2, 1, 5, 3, 10, 7}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				tem := a[i]
				a[i] = a[j]
				a[j] = tem
			}
		}
	}
	fmt.Println(a)

}
