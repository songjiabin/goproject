package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel 中读取内容
	// 新建协成
	go func() {
		//不停的读
		//当里面没有数据的时候就要停止了
		for i := 0; i < 8; i++ {
			time.Sleep(time.Second*3)
			num := <-ch
			fmt.Println(num)
		}
		//写入数据，表示已经读取完成了
		quit <- true

	}()

	//生产者 产生数字并写入channel
	producter(ch, quit)

}

//生产者   (ch)只是写   (quit)只读
func producter(ch chan<- int, quit <-chan bool) {
	x, y := 10, 20
	for {
		//监听channel数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
			fmt.Println("x,y", x, y)
		case flag := <-quit:
			fmt.Println(flag)
			return
		}
	}
}
