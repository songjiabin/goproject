package main

import (
	"time"
	"fmt"
)

func main() {
	//延时2秒打印一句话
	//method1()
	//method2()
	method3()
}

func method3() {
	fmt.Println(time.Now())
	f:=<-time.After(2 * time.Second)
	fmt.Println("时间到了",f)
}

func method2() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到了...")
}

func method1() {
	timer := time.NewTimer(2 * time.Second)
	//取数据
	t := <-timer.C
	fmt.Println("时间到了...", t)
}
