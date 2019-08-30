package main

import (
	"fmt"
)

func main2222() {
	a := 0.2
	b := 123
	fmt.Println(a * float64(b))

	fun99table()

	c := 1
	d := 2

	swap(c, d)
	abc(1, 2, 3, 4, 56,223)



}

func swap(c, d int) {

	c, d = d, c
	fmt.Println(c, d)


}

func abc(a ...int) {
	fmt.Println(a)
	for b,v := range a {
		fmt.Println(b,v)
	}

	fmt.Println(a[3:])
}

func fun99table() {

	/*
		1*1=1
		1*2=2 2*2=4
		1*3=3 2*3=6 3*3=9
		1*4=4 2*4=8 3*4=12 4*4=16
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
			if j == i {
				break
			}
		}
		fmt.Println()
	}

}

func demo()  {
	var a FUN1
	a=fun99table
	a()


}


type FUN1 func()

func main() {
	demo()
}