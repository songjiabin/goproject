package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	eatCh := eat()
	sleep := time.NewTimer(3 * time.Second)

	select {
	case v := <-eatCh:
		fmt.Println(v)
	case v := <-sleep.C:
		fmt.Println("睡觉了啦啊",v)
	//default:
	//	fmt.Println("回家打豆豆啦！")
	}

}

func eat() chan string {

	out := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "喊你去回家吃饭啦"
		close(out)
	}()

	return out
}
