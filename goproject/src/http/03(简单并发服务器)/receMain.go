package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {

	//监听
	listener, error := net.Listen("tcp", "192.168.2.133:2222")
	if error != nil {
		fmt.Println(error)
		return
	}

	defer listener.Close()

	//接收多个用户的请求

	for {
		con, errorAccept := listener.Accept()
		if errorAccept != nil {
			fmt.Println(errorAccept)
			return
		}

		//处理用户请求了
		go handleCon(con)

	}

}

//每接收一个新开启一个协成
func handleCon(con net.Conn) {

	defer con.Close()

	//获取客户端网络地址信息
	addr := con.RemoteAddr().String()
	fmt.Println("address:=>", addr)

	for {
		//开始读取发送过来的信息
		buf := make([]byte, 2048)
		n, error := con.Read(buf)
		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println(string(buf[:n]))

		//因为里面有一个换行符所以要n-1  用 nc 测试的时候
		//if string(buf[:n-1]) == "exit" {
		if string(buf[:n]) == "exit" {
			fmt.Println("退出了")
			return
		}

		//将读到的信息转换成大写 并返回给 客户端
		con.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}

}
