package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {

	type1()

	type2()

}

func type1() {
	a := make([]interface{}, 10)
	a[0] = 1
	a[2] = "基本密码"
	a[1] = Student{name: "宋佳宾", id: 23}
	for _, v := range a {
		//v是int类型吗
		//value:如果是则返回值
		//OK  :是不是该类型
		value, ok := v.(int)
		fmt.Println(value, ok)
	}
	//a里面什么类型都有 想知道都是你什么类型的
}

func type2() {
	a := make([]interface{}, 10)
	a[0] = 1
	a[2] = "基本密码"
	a[1] = Student{name: "宋佳宾", id: 23}

	for _, v := range a {
		switch value := v.(type) {
		case int:
			fmt.Printf("值：%d，类型是int\n", value, )
		case string:
			fmt.Printf("值：%d，类型是string\n", value)
		case Student:
			fmt.Printf("值：%d，类型是Student\n", value)
		}
	}
}
