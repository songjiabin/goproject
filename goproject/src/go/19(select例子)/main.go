package main

import "fmt"

var (
	//定义了两个通道
	chRead  = make(chan int, 1)
	chWrite = make(chan int, 1)
	y       = 12
)

func main() {
	chRead <- 133
	method()
}

func method() {
	select {
	case x := <-chRead:
		// chRead 中没有数据，所以case x := <-chRead读不到数据，所以这个case不能执行。
		fmt.Println("读到了数据是：", x)
	case chWrite <- y:
		//  writeCh是带缓冲区的通道，它里面是空的，可以写入1个数据，所以case writeCh <- y可以执行。
		fmt.Println("写入了数据", y)
	default:
		fmt.Println("do what you to do")
	}
}
