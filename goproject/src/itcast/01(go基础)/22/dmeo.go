package main

import (
	"fmt"
)

type name struct {
	name string
}

func main() {

	n := name{name: "宋佳宾"}
	a(&n)


}

func a(n *name)  {
	fmt.Println(&n)
}