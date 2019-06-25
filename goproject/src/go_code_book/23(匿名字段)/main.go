package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
	sex  byte
}

type Student struct {
	// 只有类型 没有名字、匿名字段、继承了Person的成员
	Person
	id      int
	address string
}

func main() {

	s := Student{
		id:      1,
		address: "北京邮电大学",
		Person: Person{
			age:  25,
			name: "宋佳宾",
			sex:  'Y',
		},
	}

	fmt.Printf("%c", s.sex)

}
