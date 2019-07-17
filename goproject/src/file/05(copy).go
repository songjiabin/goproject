package main

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

//复制一个文件
func main() {

	/**
		拷贝文件
	 */
	//srcFile := "./d.txt"
	//destFile := "./e.txt"

	//拷贝图片
	srcFile := "C:/Users/songjiabin1/Desktop/图床/abc.jpg"
	destFile := "my.jpg"

	CopyFile(srcFile, destFile)
	//CopyFile2(srcFile, destFile)
	//CopyFild3(srcFile,destFile)
}

//copy文件 使用io.copy
func CopyFile(srcFilePath, desFilePath string) {

	//打开需要读取的文件
	scrFile, e := os.Open(srcFilePath)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer scrFile.Close()

	//打开目标文件
	desFile, e := os.OpenFile(desFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer desFile.Close()

	//读取数据
	buf := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := scrFile.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		//一共拷贝的字节

		total += n
		//将数据写给目标文件
		desFile.Write(buf[:n])
	}

}

func CopyFile2(srcFilePath, desFilePath string) {

	//打开需要读取的文件
	scrFile, e := os.Open(srcFilePath)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer scrFile.Close()

	//打开目标文件
	desFile, e := os.OpenFile(desFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer desFile.Close()
	written, err := io.Copy(desFile, scrFile)

	fmt.Println("写入的数据为：", written, "错误为:", err)

}

//使用 ioUtils包 进行copy
func CopyFild3(srcFilePath, desFilePath string) {

	//读取文件 一次性的读取完成  避免大数据
	bytes, e := ioutil.ReadFile(srcFilePath)
	if e != nil {
		fmt.Println(e)
		return
	}

	//将读取到的数据直接复制到目标文件去

	e2 := ioutil.WriteFile(desFilePath, bytes, os.ModePerm)
	fmt.Println(e2)

}
