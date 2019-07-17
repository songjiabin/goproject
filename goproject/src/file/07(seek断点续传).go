package main

import (
	"strings"
	"fmt"
	"os"
	"io"
	"strconv"
)

func main() {

	/*
	 *  思路： 边复制  边记录复制的总量
	 */

	srcFile := "C:/Users/songjiabin1/Desktop/图床/abc.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	//创建一个临时文件的名字 用于存储复制的数量
	tempFile := destFile + "temp.txt"

	// 打开源文件
	fileSrc, e := os.Open(srcFile)
	defer fileSrc.Close()
	if e != nil {
		fmt.Println(e)
		return
	}

	fileDes, e := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer fileDes.Close()
	if e != nil {
		fmt.Println(e)
		return
	}

	//创建临时文件
	fileTemp, e := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if e != nil {
		fmt.Println(e)
		return
	}

	//① 读取临时文件的数据
	fileTemp.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n, _ := fileTemp.Read(bs)

	coutStr := string(bs[:n]) //已经读取的数量
	//转出数量
	coutInt, _ := strconv.ParseInt(coutStr, 10, 64) //10进制  最大64位

	//设置读写的位置
	fileSrc.Seek(coutInt, io.SeekStart)
	fileDes.Seek(coutInt, io.SeekStart)

	buf := make([]byte, 1024, 1024)
	total := int(coutInt) //读取的总量
	for {
		n, errs := fileSrc.Read(buf)
		if n == 0 || errs == io.EOF {
			fmt.Println("复制完成")
			//将临时文件删除掉
			remove := os.Remove(tempFile)
			fmt.Println(remove)
			fileTemp.Close()
			break
		}

		total += n
		fmt.Println(total)

		fileTemp.Seek(0, io.SeekStart)
		//记录下写入的文件大小
		fileTemp.WriteString(strconv.Itoa(total))

		//将数据写到目标文件中去
		fileDes.Write(buf[:n])

		//假装断电
		//if total > 8000 {
		//	panic("假装断电了。...")
		//
		//}

	}

}
