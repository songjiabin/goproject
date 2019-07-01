package main

import (
	"fmt"
	"time"
)

//有缓冲的通道
func main() {
	//创建一个有缓冲通道的 通道
	ch := make(chan int, 3)

	//缓冲区剩余数量的大小
	fmt.Println(len(ch))
	////缓冲区的大小
	fmt.Println(cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			ch <- 100
			fmt.Println("子协成", i)
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		<-ch
		fmt.Println("主协成", i)
	}

}
