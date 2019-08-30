package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOROOT())

	n:=runtime.GOMAXPROCS(1)
	fmt.Println("上次设置的cpu数量是：",n)
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)
	}
}
