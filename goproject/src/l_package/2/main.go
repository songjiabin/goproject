package main

import (
	"fmt"
	"unicode/utf8"
	"bytes"
	"time"
	"math/rand"
)

func main() {


	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}

	// Go语言默认使用UTF-8编码，对Unicode的支持非常好。但这也带来一个问题，也就是很多资料中提到的“获取字符串长度”的问题。内置的len()函数获取的是每个字符的UTF-8编码的长度和，而不是直接的字符数量。
	s := "其实就是rune"
	fmt.Println(len(s))                    // "16"
	fmt.Println(utf8.RuneCountInString(s)) // "8"

	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(", ")
	buffer.WriteString("world")

	fmt.Println(buffer.String())

	var arr1 = new([5]int)
	arr := arr1
	arr1[2] = 100
	fmt.Println(arr1[2], arr[2])

	var arr2 [5]int
	newarr := arr2
	arr2[2] = 100
	fmt.Println(arr2[2], newarr[2])

	ab := [5]int{1, 2, 3, 4, 5}
	t := ab[1:3:5]

	fmt.Println(t)

	fmt.Println("----------------------------------------")
	demo()

	fmt.Println("-----------------------------------------")

	demo2()

	fmt.Println("------------------------------------------")

	demo3()

	fmt.Println("--------------------------------")
	demo4()
}

func demo() {
	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d\n", len(sli), cap(sli))
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)

	var x = []int{2, 3, 5, 7, 11}
	fmt.Printf("切片x长度和容量：%d, %d\n", len(x), cap(x))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5] // a[low : high : max]  max-low的结果表示容量  high-low为长度
	fmt.Printf("切片t长度和容量：%d, %d\n", len(t), cap(t))

	// fmt.Println(t[2]) // panic ，索引不能超过切片的长度

}

func demo2() {
	x := make(map[string]string)
	x["two"] = "1"
	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	} else {
		fmt.Println("存在哦")
	}

}

func demo3() {
	data := []int{1, 2, 3}
	for i, _ := range data {
		data[i] *= 10
	}

	fmt.Println("data:", data) // 程序输出 data: [10 20 30]
}

func demo4() {

	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {

		//go func(str field) {
		//	str.print()
		//}(v)

		go v.print()

	}
	time.Sleep(3 * time.Second)
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}
