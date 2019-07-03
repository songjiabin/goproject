package main

import (
	"fmt"
	"net"
)

func main() {

	//接收方

	listener, e := net.Listen("tcp", "192.168.177.1:1111")
	if e != nil {
		fmt.Println(e)
		return
	}

	defer listener.Close()

	conn, e1 := listener.Accept()
	if e1 != nil {
		fmt.Println(e1)
		return
	}

	defer conn.Close()

	//读取文件名字
	sf := make([]byte, 1024)
	n, e2 := conn.Read(sf)
	if e2 != nil {
		fmt.Println(e2)
		return
	}

	fileName := string(sf[:n])

	fmt.Println(fileName)

}
