package main

import (
	"net"
	"fmt"
)

//用户类
type Client struct {
	C       chan string //用于发送数据的管道
	Address string      //网络地址
	Name    string      //用户名
}

func main() {

	listener, error := net.Listen("", "");
	if error != nil {
		fmt.Println(error)
		return
	}

	defer listener.Close()
	//主线程 循环等待连接
	for {
		con, error := listener.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}
		//来一个协成处理一个用户
		go handlerConn(con)
	}

}

func handlerConn(conn net.Conn) {

}
