package main

import (
	"time"
	"fmt"
)

func main() {
	//设置一个定时器 设置时间为2秒
	//2秒后往time通道写内容
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前的时间", time.Now())

	//这个里面就是两秒后的数据 在管道里面
	t := <-timer.C

	fmt.Println(t)
}
