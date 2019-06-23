package main

import (
	"fmt"
)

func main() {
	var i interface{} = 222
	j := i.(int)
	fmt.Printf("%T->%d\n", j, j)

	var x interface{} = "333"
	y := x.(string)
	fmt.Println(x, y)

	var (
		name int
		age  string
	)
	fmt.Println(name, age)

	fmt.Printf("%T\n", name, )

	a := "ahellow"
	fmt.Printf("%c", a[0])

	fmt.Println("-------------------")

	var b byte
	b = 'a'
	var c int
	c = int(b)
	fmt.Println(c)

}
