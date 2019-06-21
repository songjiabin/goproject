package main

import (
	"fmt"
)

func main() {
	 type1()

	//type2()
	//
	//
	//type3()

}
func type1() {
	fmt.Println("1")
	///所有定了的defer的会 你顺序调用
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func type2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func type3()  {

}
