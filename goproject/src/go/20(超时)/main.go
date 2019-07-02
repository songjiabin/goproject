package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协成
	go func() {
		for {
			select {
			case info := <-ch:
				fmt.Println("info", info)
			case <-time.After(3 * time.Second):
				//三秒超时了就
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("exit")
}
