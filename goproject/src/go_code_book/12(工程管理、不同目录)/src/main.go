package main

import (
	"./test"
	"fmt"
)

func init() {
	fmt.Println("init--->main")
}


func main() {
	fmt.Println(test.Test(1, 2))
}


