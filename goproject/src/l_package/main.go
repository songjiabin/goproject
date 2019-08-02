package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	str := "dab/abd./abd"
	fmt.Println(strings.Trim(str, "d"))

	//创建文件夹
	err := os.MkdirAll("./src/songjiabin/sjb", 0777)
	if err != nil {
		fmt.Print("....",err)
		return
	}
	fmt.Println("创建成功")


	title:="我是标题"
	sprintf := fmt.Sprintf("%%%s%%", title)
	fmt.Println(sprintf)

}


