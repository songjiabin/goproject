package main

import (
	"fmt"
	"time"
)

func main() {
	//新建一个携程
	go newTask()
	//延迟一秒打印
	for {
		fmt.Println("olderTask....")
		time.Sleep(time.Second)
	}

}

func newTask() {
	for {
		fmt.Println("newTask")
		time.Sleep(time.Second)
	}
}
