package main

import (
	"fmt"
)

type iA interface {
	eat()
}

type iB interface {
	iA
	drink()
}

type Student struct {
	name string
	id   int
}

func (s *Student) eat() {
	fmt.Println("eat")
}

func (s *Student) drink() {
	fmt.Println("drink")
}

func main() {

	var a iA
	var b iB

	//可以行
	a = b

	// b=a 不行

	fmt.Println(a, b)

}
