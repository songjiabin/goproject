package main

import (
	"fmt"
	"sync"
	"time"
)

//Mutex 是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。
//Mutex 类型的锁和线程无关，可以由不同的线程加锁和解锁。
func main() {
	//声明
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (G0)")
	//加锁mutex
	mutex.Lock()


	fmt.Println("The lock is locked.(G0)")
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (G%d)\n", i)
		}(i)
	}
	//休息一会,等待打印结果
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	//解锁mutex
	mutex.Unlock()

	fmt.Println("The lock is unlocked. (G0)")
	//休息一会,等待打印结果
	time.Sleep(time.Second)

	//demo2()
	//demo2()
}

func demo() {

	var (
		mutex sync.Mutex
		listX []int
	)

	for i := 0; i < 10; i++ {
		go func(x int) {
			mutex.Lock()
			listX = append(listX, i)
			mutex.Unlock()
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(listX)
}

func demo2() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock:", i)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked---------------")
	mutex.Unlock()
	wait.Wait()
}

func demo3() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}

}

func demo4() {
	var intVar int
	var wg sync.WaitGroup
	var mutex sync.RWMutex

	go func() {
		defer wg.Done()
		mutex.Lock()
		intVar = 4
		mutex.Unlock()
		fmt.Printf("first goroutine, intVar=%d\n", intVar)
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		intVar = 5
		mutex.Unlock()
		fmt.Printf("second goroutine, intVar=%d\n", intVar)
	}()

	wg.Add(2)
	wg.Wait()
	fmt.Println("end main goroutine")

}
