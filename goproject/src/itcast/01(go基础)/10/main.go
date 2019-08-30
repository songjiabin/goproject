package main

import "fmt"

type P1 struct {
	id int
}

type P2 struct {
	P1
	name string
}

func (p1 P1) SayHello() {
	fmt.Println("P1")
}

func (p1 P2) SayHello() {
	fmt.Println("P2")
}

func main() {
	p1 := P1{id: 1}
	p2 := P2{name: "宋佳宾"}
	p1.SayHello()
	p2.P1.SayHello()


}
