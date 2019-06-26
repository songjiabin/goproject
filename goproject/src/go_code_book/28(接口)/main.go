package main

import "fmt"

type Humaner interface {
	eat()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	sex byte
	id  int
}

//学生实现了 humaner的方法
func (s *Student) eat() {
	fmt.Println(s.name, s.id, "Student----吃了")
}

//老师实现了 humaner的方法
func (t *Teacher) eat() {
	fmt.Println(t.id, t.sex, "Teacher-----吃了")
}

//int类型的 MyStr
type MyStr int

func (myStr MyStr) eat() {
	fmt.Println("myStr--------吃了")
}

func main() {
	//学生吃了
	studentEat()
	//老师吃了
	TeachEat()
	//我自己也吃了
	myStrEat()
}

func studentEat() {
	var h Humaner
	s := &Student{
		id:   25,
		name: "宋佳宾",
	}
	h = s
	h.eat()
}

func TeachEat() {

	var h Humaner

	t := &Teacher{
		id:  24,
		sex: 'Y',
	}
	h = t
	h.eat()

}

func myStrEat() {

	var h Humaner

	var myStr MyStr = 23

	h = &myStr

	h.eat()
}
