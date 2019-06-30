package main

import (
	"fmt"
	"time"
)

//定义打印的方法
func Printer(str string) {
	for _, v := range str {
		fmt.Print(string(v))
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func main() {


	//定义协成
	go method1()
	go method2()

	//定义主协成的方法 不关闭
	i := 0
	for {
		i++
	}
}

func method1() {
	Printer("song")
}

func method2() {
	Printer("jia")
}
