package main

import (
	"fmt"
	"regexp"
)

//字符串的转换
func main() {
	正则1()
	正则2()
}
func 正则1() {
	buf := "aeeebc def ghi  abc tac adddedfca4c"
	//a任意字符b 进行匹配  abc √   abcd×
	// a\dc  a(数字)c
	result := regexp.MustCompile(`a\dc`)
	if result == nil {
		fmt.Println(result)
		return
	}
	//-1 代替所有的 返回切片
	r := result.FindAllStringSubmatch(buf, -1)
	fmt.Println(r)
}

//提取有效的小数
func 正则2() {

	buf := "2.213   2323  afjlafjla 4..4  2.233 55.3"

	// \. => 转译字符 .
	// /d+ 左边是数字（任意多个的）
	result := regexp.MustCompile(`\d+\.\d+`)

	if result == nil {
		fmt.Println(result)
		return
	}

	r := result.FindAllStringSubmatch(buf, -1)
	fmt.Println(r)

}
