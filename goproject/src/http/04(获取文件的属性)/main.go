package main

import (
	"os"
	"fmt"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage: xxx file")
	}
	fileName := list[1]
	//获取文件的属性
	fileInfo, error := os.Stat(fileName)
	if error != nil {
		fmt.Println(error)
		return
	}

	//找到路径下的名字
	fmt.Println(fileInfo.Name())
	//文件的大小
	fmt.Println(fileInfo.Size())

}
