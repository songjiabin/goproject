package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//发送方
func main() {

	//1、提示输入文件
	fmt.Println("请输入需要传输的文件")
	var path string
	fmt.Scan(&path)

	//获取文件名
	fileInfo, error := os.Stat(path)
	if error != nil {
		fmt.Println(error)
		return
	}

	//主动连接服务器
	conn, error1 := net.Dial("tcp", "192.168.2.133:2222")
	if error1 != nil {
		fmt.Println(error1)
		return
	}

	defer conn.Close()

	//给服务端发送文件的名字
	_, error2 := conn.Write([]byte(fileInfo.Name()))

	if error2 != nil {
		fmt.Println(error2)
		return
	}

	//等着对方的回复，如果对方回复了ok那么可以进行接下来的传输
	buf := make([]byte, 1024)
	n, error3 := conn.Read(buf)
	if error3 != nil {
		fmt.Println(error3)
		return
	}

	if string(buf[:n]) == "ok" {
		//说明接收成功
		//开始真正的发送文件的内容
		sendFile(path, conn)
	}

}

//发送文件
func sendFile(path string, conn net.Conn) {

	//首先打开文件、
	file, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()

	//读取里面的内容
	buf := make([]byte, 1024)
	for {
		//循环的发送
		n, error1 := file.Read(buf)
		if error1 != nil {
			if error1 == io.EOF {
				//说明文件读完了
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println(error1)
			}
			return
		}

		//发送给服务端
		conn.Write(buf[:n])

	}

}
