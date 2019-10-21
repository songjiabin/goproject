package main

import (
	"fmt"
	"regexp"
)

type Persons interface {
	sayHello() string
}

type p struct {
	name string
}

func (ps p) sayHello() string {
	return "2222"
}

func main() {
	var ps Persons
	p := p{name: "dadf"}
	p.sayHello()
	ps = &p
	fmt.Println(ps.sayHello())

	path := "/login"
	reg := regexp.MustCompile("(/v1/user|/login)")
	if !reg.MatchString(path) {
		fmt.Println("!reg.MatchString(path)")
		return
	}

	fmt.Println("math")


}
