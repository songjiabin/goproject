package main

import (
	"fmt"
	"net"
	os "os"
	"io"
)

func main() {

	//接收方

	listener, e := net.Listen("tcp", "192.168.2.133:2222")
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

	//如果接收文件名字成功 那么给客户端回复 ok
	conn.Write([]byte("ok"))

	//接收文件
	recvFile(fileName, conn)

}

func recvFile(fileName string, conn net.Conn) {
	//接收文件
	file, error := os.Create(fileName)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		//循环的读
		n, error := conn.Read(buf)
		if error != nil {
			if error == io.EOF {
				fmt.Println("文件读取完毕")
			}
			return
		}

		if n==0 {
			break
		}


		//将文件写入硬盘
		file.Write(buf[:n])
	}

}
