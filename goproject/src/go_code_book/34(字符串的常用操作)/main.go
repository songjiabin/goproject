package main

import (
	"fmt"
	"strings"
)

//字符串的操作

func main() {

	//是否包含某个字段
	fmt.Println(strings.Contains("hellogo", "hells"))

	// Join
	s := []string{"abc", "go", "hello"}
	result := strings.Join(s, "-")
	fmt.Println(result)

}
