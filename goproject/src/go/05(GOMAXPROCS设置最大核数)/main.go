package main

import (
	"fmt"
	"runtime"
)

func main() {

	//设置几核
	 runtime.GOMAXPROCS(1)


	for {
		go fmt.Print("go协成")
		fmt.Print("other协成")
	}
}
