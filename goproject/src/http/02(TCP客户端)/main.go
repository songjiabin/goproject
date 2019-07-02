package main

import (
	"net"
	"fmt"
)

func main() {
	//自动连接服务器
	con, error := net.Dial("tcp", "192.168.2.133:1000")
	defer con.Close()
	if error != nil {
		fmt.Println(error)
		return
	}

	//发送数据

	n, error1 := con.Write([]byte("are you ok?"))
	if error1 != nil {
		fmt.Println(error1)
		return
	}
	fmt.Println(n)

}
