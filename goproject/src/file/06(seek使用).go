package main

import (
	"os"
	"fmt"
	"io"
)

/**
 seek 可以进行读取指定的位置的数据
 */
func main() {

	file := "./d.txt"

	openFile, e := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer openFile.Close()

	//读写
	buf := []byte{1}
	openFile.Read(buf)
	fmt.Println(string(buf))

	//seek 定位光标的位置 从指定的位置读
	// io.SeekStart 光标开始的位置
	// io.SeekCurrent 光标的当前位置
	//  io.SeekEnd 光标后的位置
	openFile.Seek(4, io.SeekStart) //从第四位开始读包括第四位
	openFile.Read(buf)
	fmt.Println(string(buf))

	//从上个读取的位置 在读取3
	openFile.Seek(3, io.SeekCurrent)
	openFile.Read(buf)
	fmt.Println(string(buf))

	//将光标定位到最后
	openFile.Seek(0, io.SeekEnd)
	openFile.WriteString("AAAA")





}
