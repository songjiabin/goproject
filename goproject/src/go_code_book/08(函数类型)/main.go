package main

import "fmt"

type funType func(int, int) int

func fun1(a int, b int) int {

	return a + b
}

func main() {

	var a funType
	a=fun1
	fmt.Println(a(1,2))

}
