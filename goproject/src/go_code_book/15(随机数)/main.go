package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//设置种子，只需要一次

	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		//每次随机数生成的都是一样的 因为你设置的种子是一样的

		fmt.Println(rand.Int())
		//限制在10以内的
		fmt.Println(rand.Intn(100))
	}

}
