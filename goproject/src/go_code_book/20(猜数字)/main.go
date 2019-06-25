package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//随机产生一个四位数字
	var randNum int
	CreateNum(&randNum)
	fmt.Println(randNum)

	randSlice := make([]int, 4)
	//将生成的四位数放到切片里面
	getNum(randSlice, randNum)

	onGame(randSlice)
}

//开始游戏
func onGame(getNums []int) {
	//用户输入的数字
	var num int
	newSlice := make([]int, 4)

	for {
		for {
			fmt.Println("请输入一个四位数字")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			}
			fmt.Println("输入的不符合要求")
		}
		//当符合要求的时候
		//将用户输入的保存到切片中去
		getNum(newSlice, num)
		n := 0
		for i := 0; i < 4; i++ {
			if newSlice[i] > getNums[i] {
				fmt.Printf("输入的第%d个数字大了\n", i+1)
			} else if newSlice[i] < getNums[i] {
				fmt.Printf("输入的第%d个数字小了\n", i+1)
			} else {
				fmt.Printf("输入的第%d个数字正确\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("猜对了哦！！！")
			break
		}
	}
}

func CreateNum(num *int) {
	rand.Seed(time.Now().UnixNano())
	var thisNum int
	for {
		thisNum = rand.Intn(10000)
		if thisNum >= 1000 {
			break
		}
	}
	*num = thisNum
}

func getNum(slice []int, num int) {
	slice[0] = num / 1000
	slice[1] = (num % 1000) / 100
	slice[2] = num % 100 / 10
	slice[3] = num % 10
}
