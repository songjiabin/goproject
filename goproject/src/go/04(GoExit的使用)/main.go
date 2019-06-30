package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go test()

	for {
		//main协成
		a := 1
		time.Sleep(time.Second)
		fmt.Println(a)
	}
}

func test() {
	fmt.Println("aaaa")
	test1()
	fmt.Println("cccc")
}

func test1() {
	defer fmt.Println("dddd")
	//终止此协成
	runtime.Goexit()
	fmt.Println("eeeeee")

}
