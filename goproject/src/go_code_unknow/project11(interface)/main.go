package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

func main() {

	b := B{}

	var a interface{}
	a = &b
	if v, ok := a.(A); ok {
		fmt.Println(v.run())
	}

}

type A interface {
	run() string
}

type B struct {
}

func (this *B) run() string {
	return "1"
}
