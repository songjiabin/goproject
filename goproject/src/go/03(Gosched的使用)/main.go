package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go携程...")
		}
	}()

	for i := 0; i < 2; i++ {
		//当执行的时候，可以先让别的携程执行，别的携程执行完了。自己再执行
		runtime.Gosched()
		fmt.Println("主携程...")
	}
}
