package main

import (
	"fmt"
	"path/filepath"
	"path"
	"os"
	"err"
)

func main() {

	/*
		文件操作
		1、路径
		    相对路径 相对于当前的工程的路径
	        绝对路径 从根目录开始的路径
			. 当前目录
		    .. 上一层
	 */

	fileName1 := "E:/gocode/goproject/goproject/src/file/a.txt" //绝对路径
	fileName2 := "a.txt"                                        //相对路径
	//判断是否为绝对路径
	fmt.Println(filepath.IsAbs(fileName1)) //true
	fmt.Println(filepath.IsAbs(fileName2)) //false

	//返回绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	//获取父目录
	fmt.Println(path.Join(fileName1, ".."))
	fmt.Println(path.Join(fileName2, ".."))

	/*
	//创建目录 只能创建一层
	errMkdir := os.Mkdir("E:/gocode/goproject/goproject/src/file/newfile", os.ModePerm)
	err.PrintErr(errMkdir)
	fmt.Println("创建目录成功")

	//一次创建多层的文件目录
	errAll := os.MkdirAll("E:/gocode/goproject/goproject/src/file/aa/bb/cc", os.ModePerm)
	err.PrintErr(errAll)
	*/

	//文件的创建（绝对路径的文件）  如果文件存在会清空文件 如果没有文件 那么创建该文件
	file, e := os.Create("E:/gocode/goproject/goproject/src/file/aa/a.txt")
	err.PrintErr(e)
	fmt.Println(file)

	//文件的创建  相对路径的文件
	fileNew := "b.txt"
	create, i := os.Create(fileNew)
	err.PrintErr(i)
	fmt.Println(create)

	//打开一个文件  这种模式只能读
	open, i2 := os.Open("c.txt")
	defer open.Close()
	err.PrintErr(i2)

	byt := make([]byte, 1024)
	n, _ := open.Read(byt)
	fmt.Println("open并读取", open, string(byt[:n]))

	//打开一个文件  并可以读写
	openFile, i3 := os.OpenFile("c.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	defer openFile.Close()
	err.PrintErr(i3)

	//读
	byt2 := make([]byte, 1024)
	n2, _ := openFile.Read(byt2)
	fmt.Println("openFile并读取2", open, string(byt2[:n2]))

	//写
	openFile.WriteString("eeeeeeee")

	//删除目录 只能删除一个文件或者目录
	remove := os.Remove("E:/gocode/goproject/goproject/src/file/d.txt")
	err.PrintErr(remove)

	//删除所有的
	os.RemoveAll("path")




















}
