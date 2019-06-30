package main

import (
	"fmt"
	"io"
	"os"
)

//拷贝文件
func main() {
	list := os.Args
	if len(list) > 3 || len(list) < 3 {
		fmt.Println("usage:xxx srcpath dstpath")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName {
		fmt.Println("源文件路径和目的文件路径不能相同哦")
		return
	}

	//打开源文件 只读的方式
	srcFile, errorSrc := os.Open(srcFileName)
	if errorSrc != nil {
		fmt.Println(errorSrc)
		return
	}

	//新建目的文件
	dstFile, errorDst := os.Create(dstFileName)
	if errorDst != nil {
		fmt.Println(errorDst)
		return
	}

	//关闭文件
	defer srcFile.Close()
	defer dstFile.Close()

	//读取源文件
	//新建缓冲区
	buf := make([]byte, 1024*4)
	for {
		n, error := srcFile.Read(buf)
		if error != nil {
			if error == io.EOF { //读取完毕
				break
			}
			fmt.Println(error)
		}
		//将读取到的内容写到目标文件中去
		dstFile.Write(buf[:n])
	}

	//命令行执行； go build main.og
	//得到main.exe
	//命令写上源地址，目标地址
	//main.exe demo.txt demoCopy.txt
	// 这样就复制了demo.txt 了

}
