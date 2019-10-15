package main

import "fmt"

type Persons interface {
	sayHello() string
}

type p struct {
	name string
}

func (ps  p) sayHello() string {
	return "2222"
}

func main() {
	var ps Persons
	p := p{name: "dadf"}
	p.sayHello()
	ps = &p
	fmt.Print(ps.sayHello())

}
