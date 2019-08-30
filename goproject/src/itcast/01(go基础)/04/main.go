package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mainaaa() {

	//arr:=[5]int{1,2,3,4,5}
	//在定义时 叫元素个数
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//在数组使用时叫下标 （0~len(arr)-1）
	//arr[5]=100//数组下标越界
	//arr[-1]=20//数组下标越界
	//数组在定义后 元素个数就已经固定 不能添加
	//fmt.Println(arr)
	//数组是一个常量 不允许赋值 数组名代表整个数组
	//arr=10//err
	//数组名也可以表示数组的首地址
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])

	test()
	fmt.Println("-------------------------------------af ")
	test4()
}

func test() {
	arr := [10]int{9, 1, 5, 6, 7, 3, 10, 2, 4, 8}

	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(max)
	test2()
	test5()
}

func test2() {
	fmt.Println("------")
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//定义起始和结束下标
	start := 0
	end := len(arr) - 1

	for {
		if (start > end) {
			break
		}
		arr[start], arr[end] = arr[end], arr[start]

		start++
		end--
	}

	fmt.Println(arr)
}

//冒泡排序
func test4() {
	arr := [10]int{9, 1, 5, 6, 7, 3, 10, 2, 4, 8}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println(arr)

}

func test5() {
	//通过键盘输入一个字符数组 统计字符出现的次数
	//jjsdfywyeurihsdfghsudjsaghsjiwyesufjsdgfgsjgfdsh
	//j 6
	//s 10

	str := "jjsdfywyeurihsdfghsudjsaghsjiwyesufjsdgfgsjgfdsh"
	for _, v := range str {
		fmt.Printf("%c\n", v)

	}

}

func main() {
	var arr [10]int = [10]int{
		1, 2,
	}
	test6(&arr)
	fmt.Println(arr)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		intn := rand.Intn(100)
		fmt.Println(intn)
		fmt.Println(time.Now().UnixNano())
	}

	test7()

}

func test6(array *[10]int) {
	array[0] = 10
	fmt.Println(*array)
}

