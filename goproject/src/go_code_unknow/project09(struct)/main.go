package main

import (
	"fmt"
)

type person struct {
	name string
	Age  int
}

//匿名函数
type person2 struct {
	name string
	age  int
	contract struct {
		Phone, City string
	}
}

//匿名字段

type person3 struct {
	string
	int
	bool
}

type human struct {
	sex string
}

type women struct {
	human
	name string
}

type man struct {
	human
	age int
}

func main() {
	//type1()
	//type2()
	//type3()
	type4()
	type5()
	type6()
}

func type5() {
	a := man{
		age:   1,
		human: human{sex: "男"},
	}

	a.sex = "男的"
	fmt.Println(a.sex)
	fmt.Println(a)
}

func type4() {
	p := person3{
		"1",
		2,
		true,
	}
	fmt.Println(p)
}

func type3() {
	p := person2{
		age:  1,
		name: "sjb",
	}
	p.contract.City = "北京"
	p.contract.Phone = "1243432442"
	fmt.Println(p)
}

func type2() {
	//匿名结构
	a := &struct {
		sea string
		age int
	}{
		sea: "宋佳宾",
		age: 12,
	}
	fmt.Println(a)
}

func type1() {
	p := &person{
		name: "name",
		Age:  33,
	}
	updateAge(p)
	fmt.Println(p)
}

func updateAge(per *person) {
	per.Age = 27
}

type A struct {
	B
	name string
}

type B struct {
	name string
}

func type6() {
	a := A{
		name: "11",
		B:    B{name: "222"},
	}
	fmt.Println(a.name,a.B.name)
}
