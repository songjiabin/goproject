package main

import "fmt"

type Person interface {
	eat()
}

type Student struct {
	name string
}

func (s *Student) eat() {
	fmt.Println(s.name + "eating")
}

func sayHello(person Person) {
	person.eat()
}

func main() {
	s := Student{name: "宋佳宾"}

	var p Person

	p = &s

	p.eat()






}
