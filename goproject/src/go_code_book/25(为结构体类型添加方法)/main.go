package main

import "fmt"

type Student struct {
	name string
	age  int
	sex  byte
}

//等于是 类里面的方法
func (s Student) printInfo() {
	fmt.Println(s)
}

//因为是要修改成员里面的变量 所以要用指针操作
func (s *Student) setPersonInfo(name string, age int, sex byte) {

	s.name = name
	s.age = age
	s.sex = sex

}

func main() {

	var s Student
	//赋值操作
	(&s).setPersonInfo("宋佳宾", 10, 'Y')
	s.printInfo()

}
