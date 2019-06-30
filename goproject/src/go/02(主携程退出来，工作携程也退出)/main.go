package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println("工作携程....")
			time.Sleep(time.Second)
		}
	}()

	i := 0;
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("主携程....")
		time.Sleep(time.Second)
	}
}
