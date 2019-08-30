package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)
	file, e := os.OpenFile(path, os.O_RDONLY, os.ModeDir) //os.modedir 读取目录的权限
	if e != nil {
		fmt.Print("openFile_error", e)
		return
	}
	defer file.Close()

	//读取目录项
	infos, e := file.Readdir(-1) //-1 所有的目录
	if e != nil {
		fmt.Print("readdir_error", e)
		return
	}

	for _, v := range infos {
		if v.IsDir() { //是目录的话
			fmt.Println(v.Name(), "是一个目录")
		} else {
			fmt.Println(v.Name(), "是一个文件")
		}
	}

}
