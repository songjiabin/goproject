package main

import "fmt"

func test(a int) {

	//设置匿名函数
	defer func() {
		//可以让错误不打印
		//recover()
		if error := recover(); error != nil {
			//只有在这种情况下 产生了异常
			fmt.Println(error)
		}
	}()

	var arrays [10]int

	arrays[a] = 23

	fmt.Println("---------------------")

}

func main() {

	test(1)

}
