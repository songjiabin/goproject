package main

import (
	"os"
	"bufio"
	"fmt"
)

//bufio下写的操作
func main() {





	fileName := "C:/Users/songjiabin1/Desktop/abc.txt"
	file, e := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer file.Close()

	//先将数据写入缓冲区，再将数据写入到磁盘中。如果缓存去的大小小于写入数据的大小，直接写入磁盘。

	writer := bufio.NewWriter(file)

	i, e := writer.WriteString("abcddddddddddddddddddddddddddddddddddddddddddddddddddddd")

	//这个时候是没数据的，因为数据写入到了write中的缓冲区中
	fmt.Println("写入的大小是：", i,e)

	writer.Flush()// 将缓冲区的数据push到磁盘中

}
