package main

import (
	"fmt"
	"time"
)

//定义一个管道 全局的变量

var ch = make(chan int)

func main() {

	go method1()
	go method2()

	i := 0
	for {
		i++
	}
}

//目的是先执行方法1 在执行方法2

func method1() {
	Printer("jia")
	//在执行方法后往管道里面放数据
	ch <- 200
}

func method2() {
	//当开始执行的时候 从管道里面取数据，当发现里面没有那么阻塞,
	// 直到里面有数据了，取到了才开始执行
	<-ch
	Printer("song")

}

func Printer(str string) {
	for _, value := range str {
		fmt.Print(string(value))
		time.Sleep(time.Second)
	}
	fmt.Println()
}
