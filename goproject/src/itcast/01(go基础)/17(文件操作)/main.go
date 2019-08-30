package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//create()
	//open()
	//write()
	//openFile()
	//bufIo()

	bigFileCopy()
}

func create() {
	file, e := os.Create("C:/Users/songjiabin1/Desktop/MyShare/demo.txt") //如果没有重新创建，如果有覆盖
	if e != nil {
		fmt.Println("err-->", e)
	}
	defer file.Close()
	fmt.Println("成功 ")
}

func open() {
	file, e := os.Open("C:/Users/songjiabin1/Desktop/MyShare/demo.txt")
	if e != nil {
		fmt.Println("error-->", e)
	}
	defer file.Close()

}

func write() {
	//这样写不了的，因为open的权限是只读。没有权限
	file, e := os.Open("C:/Users/songjiabin1/Desktop/MyShare/demo.txt")
	if e != nil {
		fmt.Println("err-->", e)
	}
	defer file.Close()

	_, _ = file.WriteString("你是说呢？")

}

func openFile() {
	//打开文件并且方式是 读和写
	file, e := os.OpenFile("C:/Users/songjiabin1/Desktop/MyShare/demo.txt", os.O_RDWR, os.ModePerm)
	if e != nil {
		fmt.Println("err--->", e)
	}
	defer file.Close()
	file.WriteString("#####")

}

//bufio操作
func bufIo() {
	file, e := os.OpenFile("C:/Users/songjiabin1/Desktop/MyShare/demo.txt", os.O_RDWR, os.ModePerm)
	if e != nil {
		fmt.Println("err--->", e)
	}
	defer file.Close()
	fmt.Println("successful")
	//创建一个带有缓冲区的reader
	reader := bufio.NewReader(file)

	for {
		bytes, e := reader.ReadBytes('\n') //以换行符区分
		if e != nil && e == io.EOF {

			break
		} else if e != nil {
			fmt.Println("err--->", e)
		}

		fmt.Println(string(bytes))
	}

}

//大文件拷贝
func bigFileCopy() {
	srcPath := "C:/Users/songjiabin1/Desktop/MyShare/1.复习.ev4"
	srcFile, e := os.OpenFile(srcPath, os.O_RDWR, os.ModePerm)
	if e != nil {
		fmt.Println("srcError->", e)
		return
	}
	defer srcFile.Close()

	createPath := "C:/Users/songjiabin1/Desktop/copy/加密.ev4"
	createFile, e := os.Create(createPath)
	if e != nil {
		fmt.Println("createFile-->", e)
		return
	}
	defer createFile.Close()

	desPath, e := os.OpenFile(createPath,os.O_RDWR,os.ModePerm)
	if e != nil {
		fmt.Println("desPathFile-->", e)
		return
	}
	defer desPath.Close()

	//从scrFile文件中读取数据到缓冲区中
	buf := make([]byte, 1024*4)
	for {
		n, e := srcFile.Read(buf)
		if e != nil && e == io.EOF {
			break
		} else {
			desPath.Write(buf[:n])
		}
	}

	fmt.Println("复制完成")

}


