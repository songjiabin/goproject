package main

import "fmt"

type pointer int

//接受者的类型不能为指针类型 如:
//type pointer *int

func (p *pointer) test() {
	fmt.Println(*p)
}

func main() {

	var p pointer
	p = 2
	(&p).test()

}
