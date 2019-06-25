package main

import (
	"fmt"
)

//定义一个结构体类型

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {

	func1()

	func2()

	fun3()

}

func fun3() {

}

//指针类型的结构体创建
func func2() {
	//结构体 成员变量的使用 -指针变量
	var s3 Student
	p := &s3
	p.name = "宋佳宾3"
	p.id = 123
	p.age = 23
	fmt.Println(*p)
	//第二个方式
	// 开辟一个快内存 p2指向改内存地址
	p2 := new(Student)
	p2.name = "许美猴"
	fmt.Println(p2)
}

//结构体的声明
func func1() {
	var student Student = Student{id: 1}
	fmt.Println(student.id)
	s := &Student{
		id:   1,
		name: "宋佳宾",
	}
	fmt.Println(*s)
	//指针类型的
	var s2 *Student = &Student{
		id:   3,
		name: "宋佳宾2",
	}
	fmt.Println(*s2)
}
