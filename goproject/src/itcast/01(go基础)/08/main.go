package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	score int
	name string
}




func main() {

	s := Student{Person: Person{name: "宋佳宾", age: 23}, score: 23,name:"学生"}
	fmt.Println(s.Person.name)
	fmt.Println(s.score)
	fmt.Println(s.name)
}
