package main

import (
	"fmt"
	"time"

)

func producer(ch chan int, d time.Duration, num int) {

	for i := 0; i < num; i++ {
		ch <- i
		time.Sleep(d)

	}
}

func main() {


	ch := make(chan int)

	go producer(ch, 100*time.Millisecond, 2)
	go producer(ch, 200*time.Millisecond, 5)

	//for {
	//	fmt.Println(<-ch)
	//}

	for {
		select {
		case info := <-ch:
			fmt.Println(info)
		}
	}

	defer close(ch)
}
