package main

import "fmt"

type Person struct {
	age  int
	name string
	sex  byte
}

type Student struct {
	*Person
	address string
	id      int
}

//结构体指针类型匿名字段
func main() {

	//创建方式1 
	student := Student{
		address: "北京邮电大学",
		id:      25,
		Person: &Person{
			name: "宋佳宾",
			age:  25,
			sex:  'Y',
		},
	}

   fmt.Println(student)
	
	
	
	//创建方式2
	student2:=Student{
		address: "北京邮电大学",
		id:      25,
		Person:new(Person),
	}

	student2.Person.age=1
	student2.Person.name="宋佳宾"
	student.Person.sex='Y'


	fmt.Println(*student2.Person)

	
	
	
}
