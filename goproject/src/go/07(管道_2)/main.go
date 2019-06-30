package main

import (
	"fmt"
	"time"
)

//实现管道
func main() {

	defer fmt.Println("主协成执行完毕")

	ch := make(chan string)
	//让字协成完成后 再关闭主协成
	go func() {
		defer fmt.Println("子协成执行完毕")
		for i := 0; i < 3; i++ {
			fmt.Println("子协成进行打印中....", i)
			time.Sleep(time.Second)
		}
		ch <- "宋"
	}()

	str, _ := <-ch
	fmt.Println(str)
}
