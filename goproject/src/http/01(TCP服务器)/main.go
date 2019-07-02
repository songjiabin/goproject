package main

import (
	"net"
	"fmt"
)

func main() {

	//1、监听等待用户的连接
	listener, error := net.Listen("tcp", "192.168.2.133:1000")
	if error != nil {
		fmt.Println(error)
		return
	}
	defer listener.Close()

	//阻塞用户连接
	for {
		//循环接收
		con, error := listener.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}

		buf := make([]byte, 1024)
		//接收用户的请求
		n, error2 := con.Read(buf)
		//读了多少个n
		if error2 != nil {
			fmt.Println(error2)
		}

		fmt.Println("buf", string(buf[:n]))

		defer con.Close()
	}
}
