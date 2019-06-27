package main

import (
	"fmt"
)

//定义接口
type Human interface {
	eat()
}

//定义接口
type Person interface {
	Human
	drink()
}

//学生
type Student struct {
	name string
	age  int
}

func (s *Student) eat() {
	fmt.Println("Student eat ....")
}

func (p *Student) drink() {
	fmt.Println("Student drink")
}

func main() {

	s := &Student{
		name: "宋佳",
		age:  23,
	}

	s.drink()
	s.eat()

	var human Human
	var person Person

	human = s
	person = s

	human.eat()
	person.drink()

}
