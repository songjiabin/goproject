package main

import (
	"fmt"
	"time"
)

//实现管道
func main() {

	//demo()
	demo1()
}
func demo() {
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

func demo1() {

	ch := make(chan string)
	defer fmt.Println("主协成执行完毕")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协成...", i)
		}
		defer fmt.Println("子协成执行完毕")
		ch <- "交给你执行了 主协成"
	}()

	i := <-ch
	fmt.Println(i)

}
