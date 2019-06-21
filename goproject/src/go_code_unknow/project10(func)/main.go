package main

import "fmt"

type A struct {
	name string
}

type B struct {
	name string
}

func (a *A) Prict(name string) {
	a.name = "宋佳宾"

}

func (b B) Prict(name string) {
	fmt.Println(b)
}

func main() {
	//a := &A{
	//	name: "名字",
	//}
	//a.Prict(a.name)
	//fmt.Println(a)
	//
	//b := B{
	//	name: "名字B",
	//}
	//b.Prict(b.name)

	str := "song"
	for i, j := 0, len(str); i < j; i++ {

		 fmt.Println(string(str[i]))
	}

}
