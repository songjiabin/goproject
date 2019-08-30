package main

import "fmt"

func main() {

	a, b := 1, 2

	a, b = b, a

	fmt.Println(a, b)

	x, y := 1, 2
	swap(&x,&y)
	fmt.Println(x,y)
}

func swap(x, y *int) {
	 x,  y =  y,  x
}
