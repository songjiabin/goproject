package main

import (
	"os"
	"err"
	"fmt"
)

func main() {

	info, e := os.Stat("E:/gocode/goproject/goproject/src/file/a.txt")
	err.PrintErr(e)
	fmt.Printf("类型是：%T\n", info)
	//文件的名字
	fmt.Println(info.Name())
	//文件的大小
	fmt.Println(info.Size())
	//改路径是否为目录
	fmt.Println(info.IsDir())
	//改文件的时间 最后一次的修改时间
	fmt.Println(info.ModTime())
	//改文件的操作权限
	//-rw-rw-rw-
	fmt.Println(info.Mode())

}
