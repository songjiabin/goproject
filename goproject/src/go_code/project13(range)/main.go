package main

import (
	"fmt"

)

// for 循环的 range 形式可遍历切片或映射。
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	type1()
	type2()
}

func type1() {
	for i, v := range []int{2, 3, 5, 5} {
		fmt.Printf("当前值的下标是：%v,当前的值是：%v \n", i, v)
	}
}

func type2() {
	//创建长度是10的数组
	pow := make([]int, 10)

	for i := range pow {
		//添加数据
		pow[i] = i << uint(i) // == 2**i
	}

	//打印值
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

}
