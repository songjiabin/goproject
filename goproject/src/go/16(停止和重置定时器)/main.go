package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	go func() {
		t := <-timer.C
		fmt.Println(t)

		//重置两秒
		timer.Reset(2 * time.Second)
		t2 := <-timer.C
		fmt.Println(t2)

	}()

	//timer.Stop()
	i := 0
	for {
		i++
	}
}
