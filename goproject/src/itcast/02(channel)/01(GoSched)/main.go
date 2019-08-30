package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.Gosched()  出让当前go程所占用的 cpu时间片。当再次获得cpu时，从出让位置继续回复执行。

	go func() {
		for {
			fmt.Println("this is test fun")
		}
	}()

	for {
		fmt.Println("this is main fun")
		runtime.Gosched()
	}

}
