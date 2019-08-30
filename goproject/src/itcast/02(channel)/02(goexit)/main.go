package main

import (
	"fmt"
	"runtime"
)

//  goExit 退出当前go程序
func main() {
	go test()
	for {
		;
	}
}

func test() {

	fmt.Println("aaaaaaaaaa")
	runtime.Goexit()
	defer  fmt.Println("bbbbbbbbbbbbb")

}
