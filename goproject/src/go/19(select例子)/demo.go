package main

import (
	"time"
	"math/rand"
	"fmt"
)

func eats() chan string {

	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "喊你回家吃饭了..."
	}()
	return out
}

func main() {
	eat := eats()
	t := time.NewTimer(3 * time.Second)
	select {
	case info := <-eat:
		fmt.Println(info)
	case info := <-t.C:
		fmt.Println("到了暑假的时间了哦", info)
	//default:
	//	fmt.Println("其他的时间...")
	}
}
