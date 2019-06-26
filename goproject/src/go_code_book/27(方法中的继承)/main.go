package main

import "fmt"

/** 方法的继承
 */

type Person struct {
	name string
	age  int
	sex  byte
}

func (p *Person) printInfo() {
	fmt.Printf("%s,%d,%c", p.name, p.age, p.sex)
	fmt.Println("\n")
}

type Student struct {
	Person
	id      int
	address string
}

func main() {

	s := Student{
		Person{
			name: "宋佳宾",
			age:  23,
			sex:  'Y',
		},
		20,
		"北京邮电大学",
	}
	s.printInfo()

	//定义为一个方法
	fun := s.printInfo
	fun()
	fmt.Println("---------------------------------------------")
	f := (*Student).printInfo
	f(&s)

}
