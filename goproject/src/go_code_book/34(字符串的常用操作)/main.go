package main

import (
	"fmt"
	"strings"
)

//字符串的操作

func main() {

	//是否包含某个字段
	fmt.Println(strings.Contains("hellogo", "hells"))

	// Join 组合
	s := []string{"abc", "go", "hello"}
	result := strings.Join(s, "-")
	fmt.Println(result)

	// index  开始的位置
	fmt.Println(strings.Index("abchello", "hello"))
	//不存在的话返回 -1
	fmt.Println(strings.Index("abchello", "ahello"))

	//重复几次 进行创建
	str := strings.Repeat("SongJiaBin", 3)
	fmt.Println(str)

	// split
	newstr := strings.Split("hello", "e") //->[h llo]
	fmt.Println(newstr)

	// Trim 去掉两头的字符
	a := "#####你是谁呢#####"
	fmt.Println(strings.Trim(a, "#"))

	// 已空格进行分割，放到一个切片中去
	myStr := " are you ok?"
	for _, data := range strings.Fields(myStr) {
		fmt.Println(data)
	}



	fmt.Println(strings.Count("a bb","a"))






}
