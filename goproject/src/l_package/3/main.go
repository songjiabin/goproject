package main

import (
	"fmt"
)

func main() {

	f, e := Sqrt(-10)
	fmt.Println(f, e)
	fmt.Println("00000000000000000000000000000000000000000")
	// 捕捉内部的Panic错误
	div(10, 0)
	fmt.Println("到这里了。。。")
	div(10, -1)
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("square root of negative number %g", f)
	} else {
		return 0, nil
	}
}

func div(a, b int) {

	defer func() {

		if i := recover(); i != nil {
			fmt.Println("捕捉到了错误：", i)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)
}


func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return 0// 必须要有终止语句，如果这里没有return，则会编译错误：missing return at end of function
}