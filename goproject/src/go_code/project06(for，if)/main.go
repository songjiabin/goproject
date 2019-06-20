package main

import (
	"fmt"
	"math"
)

// for 循环很别扭 
func main() {
	//type1()
	type2()
	type2()
	type3()
	type4()
}
func type1() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	//无限循环的 for
	sum := 1
	for ; sum < 100; {
		fmt.Println(sum)
		sum = sum + 1;
	}
	fmt.Println(sum)
	// 在go中的 for 就是while
	par := 100
	for par < 1000 {
		par = par + 100
		fmt.Println(par)
	}
	//无限循环
	for {
		fmt.Println("-------------------")
		break
	}
	//if   ()不是必须的  {}是必须的
	if (2 > 1) {
		fmt.Println("if语句")
	}
	if p := math.Pow(2, 2); p > 10 {
		fmt.Println("p>10")
	} else {
		fmt.Println("P>10")
	}
	if demo := 1; demo < 2 {
		fmt.Println(demo, "if语句判断")
	}
}

//无线循环的for
func type2() {
	a := 0
	for {
		a++;
		fmt.Println(111)
		if a > 10 {
			break
		}
	}
}

//switch
func type3() {
	a := 1
	switch a {
	case 1:
		// 如果到1执行完成了  还可以继续向下执行
		fallthrough
	case 2:

	}
}

//标签的使用
func type4() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//当执行到这里的时候 将break到 LABEL1
				break LABEL1
			}
		}
	}

	fmt.Println("OK")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//当执行到这里的时候 将break到 LABEL1
				goto LABEL2
			}
		}
	}

LABEL2:
	fmt.Println("OK2")
}
