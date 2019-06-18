package main

import (
	"fmt"
	_ "math"
)

func main() {

	//hypot 位函数

hypot := func(x, y float64) float64 {
	return x + y
}

//fmt.Println(hypot(1, 2))

fmt.Println(compute(hypot))

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
