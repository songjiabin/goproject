package main

import (
	"fmt"
	"math/rand"
	"time"

)

func initData(data []int) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		data[i] = rand.Intn(100)
	}
}

//排序
func sort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
}

func main() {

	// 创建一个切片
	data := make([]int, 10)
	//初始化数据
	initData(data)
	//排序
	sort(data)

	fmt.Println(data, cap(data))
}
