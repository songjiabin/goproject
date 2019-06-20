package main

import (
	"fmt"
	"math"
)

// for 循环很别扭 
func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	//无限循环的 for 
	sum := 1
	for ; sum < 100; {
		fmt.Println(sum)
		sum = sum + 1;
	}
	fmt.Println(sum)

	// 在go中的 for 就是while

	par := 100
	for par < 1000 {
		par = par + 100
		fmt.Println(par)
	}

	//无限循环
	for {
		fmt.Println("-------------------")
		break
	}

	//if   ()不是必须的  {}是必须的
	if (2 > 1) {
		fmt.Println("if语句")
	}

	if p := math.Pow(2, 2); p > 10 {
		fmt.Println("p>10")
	} else {
		fmt.Println("P>10")
	}

	if demo := 1; demo < 2 {
		fmt.Println(demo,"if语句判断")
	}

}
