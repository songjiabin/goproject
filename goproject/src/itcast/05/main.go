package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//test1()
	test2()
}

func test1() {

	arr := [20]byte{
		's',
		'd',
		'j',
		'f',
		'j',
		's',
		'd',
		'f',
		'u',
		'i',
		'y',
		'e',
		'u',
		'd',
		'h',
		'a',
		'j',
		's',
		'h',
	}
	var ch [26]int //用来统计字符个数
	for i := 0; i < len(arr); i++ {
		ch[arr[i]-'a']++
	}

	//打印字母出现次数
	for i := 0; i < len(ch); i++ {
		if ch[i] > 0 {
			fmt.Printf("字母：%c出现：%d次\n", 'a'+i, ch[i])
		}
	}

	fmt.Println("22222")
}

//双色球
func test2() {
	//获取随机数种子
	rand.Seed(time.Now().UnixNano())

	var redball [6]int

	for i := 0; i < len(redball); i++ {
		//遍历之前存在的值和新随机数是否有重复
		//redball[i] = rand.Intn(33)+1//0-32
		temp := rand.Intn(33) + 1
		for j := 0; j < i; j++ {
			if temp == redball[j] {
				temp = rand.Intn(33) + 1
				j = 0
				continue
			}
		}

		redball[i] = temp
	}

	fmt.Println(redball, "+", rand.Intn(16)+1)
}
