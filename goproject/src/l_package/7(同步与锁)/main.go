package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	///任务完成
	wg := sync.WaitGroup{}

	//互斥锁
	var mutex sync.Mutex
	fmt.Println("Locking  (G0)")
	// 获得锁，
	mutex.Lock()
	fmt.Println("locked (G0)")
	//添加三个任务到任务池中
	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")

	wg.Wait()
}