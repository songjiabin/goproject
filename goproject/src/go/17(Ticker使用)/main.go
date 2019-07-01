package main

import (
	"time"
	"fmt"
)

func main() {
	//每隔一秒钟向通道发送一个消息
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println("i==>", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}

}
