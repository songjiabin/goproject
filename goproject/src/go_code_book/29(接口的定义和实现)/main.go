package main

import (
	"fmt"
)

type eatFood interface {
	eat()
}

type Student struct {
	name string
	age  int
}

type Person struct {
	name string
	age  int
	sex  byte
}

type myStr int

func (s *Student) eat() {
	fmt.Println("Student eat ....")
}

func (p *Person) eat() {
	fmt.Println("Person eat ...")
}

func (str myStr) eat() {
	fmt.Println("myStr eat ....")
}

//多肽的实现
func test(food eatFood) {
	food.eat()
}

//多态的实现
func main() {

	student := &Student{name: "宋佳", age: 23}

	person := &Person{name: "宋佳_person", age: 23}

	var mystr myStr = 23

	test(student)
	test(person)
	test(&mystr)

	fmt.Println("#########################################")

	//创建一个切片
	eatfood := make([]eatFood, 3)
	eatfood[0] = student
	eatfood[1] = person
	eatfood[2] = &mystr

	for _, v := range eatfood {
		v.eat()
	}
}
