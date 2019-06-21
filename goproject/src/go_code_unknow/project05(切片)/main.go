package main

import "fmt"

//切片 Slice
func main() {

	//声明切片
	var s1 []int
	fmt.Println(s1)

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Println(string(sa))

}
