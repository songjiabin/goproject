package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//客户端 主动连接服务器
	con, error := net.Dial("tcp", "192.168.2.133:2222")
	if error != nil {
		fmt.Println(error)
		return
	}

	defer con.Close()

	//接收服务器返回过来的数据
	go func() {
		//切片用了缓存数据
		buf := make([]byte, 1024*2)
		for {
			n, error := con.Read(buf)
			if error != nil {
				fmt.Println(error)
				return
			}
			//处理读到的数据
			fmt.Println("读到的数据是：", string(buf[:n]))
		}
	}()

	str := make([]byte, 1024)
	//从键盘输入内容 给服务器
	for {
		//从键盘读取内容
		n, error := os.Stdin.Read(str)
		if error != nil {
			fmt.Println(error)
			return
		}
		//将读到的内容发送给服务器
		con.Write(str[:n])

	}

}
