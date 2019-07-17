package main

import (
	"os"
	"fmt"
)

func main() {

	filePath := "./d.txt"

	//因为第一次的时候 d.txt是没有的，所以这里要使用 .openFile()设置模式
	// os.O_CREATE 没有的话可以创建
	// os.os.O_WRONLY 有写的权限
	// os.O_APPEND 进行追加的
	// os.ModePerm 权限
	file, e := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if e != nil {
		fmt.Println(e)
		return
	}

	defer file.Close()

	//写数据
	bs := []byte{65, 66, 67, 68, 69, 70} //ABCDEF
	n, err := file.Write(bs)
	fmt.Println(n)
	fmt.Println(err)

}
