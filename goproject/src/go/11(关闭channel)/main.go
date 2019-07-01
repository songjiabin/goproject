package main

import (
	"fmt"
)

func main() {
	//创建一个无缓冲的 channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			//往通道写数据
			ch <- i
		}
		//不需要再写数据了 ，关闭 channel
		close(ch)
	}()

	//使用for迭代
	/*for {
		if i, ok := <-ch; ok {
			//如果ok为true,说明管道没有关闭
			fmt.Println("读数据...", i)
		}else {
			break
		}
	}*/

	for num := range ch {
		fmt.Println(num)
	}

	defer fmt.Println("完成")

}
