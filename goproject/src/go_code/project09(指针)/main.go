package main

import (
	"fmt"
)

type abc int;

// 通 & 来 取变量值

// 通 * 通过指针间接访问目标对象

func main() {
	//定义 int a
	a := 200
	// var b *int =&a;//定义int型指针  并将a的内存地址复制给他
	b := &a
	*b++
	fmt.Println(a)

	point1()

}

//练习指针1
func point1() {

	//定义变量 i,j
	i, j := 42, 2701
	//
	p := &i; //p等于i的内存地址

	fmt.Println(*p) //根据p里面的内存地址从而得到i的值

	*p = 21
	fmt.Println(i) //通过指针的值来设定i的值

	p = &j       //p等于j的内存地址
	*p = *p / 37 //通过指针对j进行除法运算
	fmt.Println(j)

}
