package main

import (
	"fmt"
	"sync"
)

var x2 = 0

func increments(wg *sync.WaitGroup, m chan bool) {

	m <- true
	x2 = x2 + 1
	<-m
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	//通道里面是
	var m = make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increments(&w, m)
	}
	w.Wait()
	fmt.Println("final value of x", x2)
}
