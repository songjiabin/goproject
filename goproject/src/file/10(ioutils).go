package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"

)

// 使用ioutils包

func main() {

	//使用 raed
	//readMehtod()

	//使用write
	//writeMehtod()

	//使用readall
	//readAllMethod()

	//readDir 读取目录  可以读取  但是只能读取一层
	//readDir()

	//临时目录和临时文件
	//创建一个临时目录  在指定的目录下面创建一个临时文件
	tmepDirAndFile()

}
func tmepDirAndFile() {

	//创建临时的目录
	dir, err := ioutil.TempDir("C:/Users/songjiabin1/Desktop", "Test")
	//当不用的时候记得删除

	fmt.Println(dir, err)
	//创建的临时的文件
	f, err := ioutil.TempFile(dir, "Test")
	fmt.Println(f.Name())

	remove := os.Remove(dir)
	fmt.Println(remove)
	e := os.Remove(f.Name())
	fmt.Println(e)

}

func readDir() {
	fileDir := "C:/Users/songjiabin1/Desktop/图床"
	infos, e := ioutil.ReadDir(fileDir)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(len(infos))
	for _, v := range infos {
		fmt.Println(v.Name(), v.IsDir())
	}
}

func readAllMethod() {
	fileName := "C:/Users/songjiabin1/Desktop/经销存.txt"
	readerFile := strings.NewReader(fileName)
	bytes, e := ioutil.ReadAll(readerFile)
	fmt.Println(e)
	fmt.Println(string(bytes))
}

func readMehtod() {
	//读取文件中的所有的数据
	fileName := "C:/Users/songjiabin1/Desktop/经销存.txt"
	content, e := ioutil.ReadFile(fileName)
	if e == nil {
		fmt.Println("读取完成...", string(content))
		return
	}
}

func writeMehtod() {
	//写入数据
	fileNameWrite := "C:/Users/songjiabin1/Desktop/cccccc.txt"
	contentWrite := "hello world"
	fileError := ioutil.WriteFile(fileNameWrite, []byte(contentWrite), os.ModePerm)
	//if fileError != nil {
	//	fmt.Println("写入完成...")
	//	return
	//}
	fmt.Println(fileError)
}
