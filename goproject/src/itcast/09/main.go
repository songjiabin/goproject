package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	*person
	*Student
	id    int
	score int
}


func (p *person)  test()  {
	fmt.Println("test方法",p.name,p.age,p.sex)
}



func main() {

	s := Student{person: &person{name: "宋佳宾", sex: "男", age: 23}, id: 223, score: 23}
	fmt.Println(s.person.name)
	fmt.Println(s.id)
	fmt.Println(s.score)
	s.person = new(person)
	s.person.age = 1
	s.person.sex = "女"
	s.person.name = "new"
	s.Student = &Student{score: 23}
	fmt.Println(s.person)
	fmt.Println(s.Student.score)


	s.test()


}
