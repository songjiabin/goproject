package main

import "fmt"

func main() {
	//创建一个双向的
	ch := make(chan int)
	//生产者、生成数字、写入channel
	go producter(ch)
	//消费者，从channel读取内容，打印
	consummer(ch)
}

//新开一个协成
//此通道只能写不能读
func producter(out chan<- int) {
	for i := 0; i < 10; i++ {
		//通道里面写数据
		out <- i * i
	}

	close(out)
}

//此通道只能读不能写
func consummer(in <-chan int) {
	//读取
	for num := range in {
		fmt.Println("num=", num)
	}
}
