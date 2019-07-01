package main

import (
	"fmt"
	"time"
)

//无缓冲的channel 就是管道里面是数量必须为0  当管道里面没有数据的时候才可以存入数据
func main() {

	//创建一个无缓冲的channel
	ch := make(chan int)

	//缓冲区剩余数量的大小
	fmt.Println(len(ch))
	////缓冲区的大小
	fmt.Println(cap(ch))

	//新建协成
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子协成....", i)
			ch <- 10
		}
	}()

	time.Sleep(2*time.Second)
	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("主协成....", i)
	}

}
