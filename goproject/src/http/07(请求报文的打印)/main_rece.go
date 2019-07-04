package main

import (
	"net"
	"fmt"
)

func main() {

	listener, error := net.Listen("tcp", "192.168.56.1:2122")
	if error != nil {
		fmt.Println(error)
		return
	}

	conn, error_accept := listener.Accept()
	if error_accept != nil {
		fmt.Println(error_accept)
		return
	}

	//接受接受到的消息
	buf := make([]byte, 1024)
	n, error_read := conn.Read(buf)
	if error_read != nil {
		fmt.Println(error_read)
		return
	}

	fmt.Println(string(buf[:n]))

	//用浏览器 访问 则会出现报文
	a := `GET / HTTP/1.1  请求行
          Host: 192.168.56.1:2222
          User-Agent: Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36
          Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3
          Accept-Encoding: gzip, deflate
          Accept-Language: zh-CN,zh;q=0.9
          Cache-Control: max-age=0
          Proxy-Connection: keep-alive
          Upgrade-Insecure-Requests: 1
          X-Lantern-Version: 5.4.7
    	  `

	fmt.Println(a)

}
