package main


import (
	"math"
	"fmt"
)

func main() {
	//向上取整
	ceil := math.Ceil(1.1)
	fmt.Println(ceil)

	//向下取整
	floor := math.Floor(1.1)
	fmt.Println(floor)
}
