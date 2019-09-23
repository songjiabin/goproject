package main

import (
	"fmt"
	"math"
)

func main() {
	//向上取整
	ceil := math.Ceil(1.1)
	fmt.Println(ceil)

	//向下取整
	floor := math.Floor(1.1)
	fmt.Println(floor)

	//if preparer, ok := this.AppController.(NestPreparer); ok {
	//		preparer.NextPreparse()
	//	}

	var a Abc;
	a.run()


}

type Abc interface {
	run() string
}

type person struct {
}

func (this *person) run() string {
	return "ss"
}
