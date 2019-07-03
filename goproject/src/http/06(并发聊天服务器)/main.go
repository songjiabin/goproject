package main

import (
	"net"
	"fmt"
	"time"
)

//用户类
type Client struct {
	C       chan string //用户发送数据的管道
	Address string      //网络地址
	Name    string      //用户名
}

//保存在线的用户
var onlineMap map[string]Client

//管道，当有新消息了  则通过此管道通知
var message = make(chan string)

func main() {

	listener, error := net.Listen("tcp", "192.168.177.1:2222");
	if error != nil {
		fmt.Println(error)
		return
	}

	defer listener.Close()

	//新建一个协成，转发信息，只要有消息了就遍历map。给map每个成员发送此消息
	go Manager()

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

//处理连接
func handlerConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的地址
	addr := conn.RemoteAddr().String()

	//结构体
	cli := Client{
		Address: addr,
		Name:    addr,
		C:       make(chan string),
	}

	onlineMap[addr] = cli

	//新建一个协成  专门给当前的客户发送信息
	go writeInfoToClient(cli, conn)

	message <- "...." + cli.Address

	time.Sleep(200000000 * time.Second)

	defer fmt.Println("结束")

}

func writeInfoToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		fmt.Println("发送信息", client.Address)
		conn.Write([]byte(msg + "\n"))
	}
	//select {
	//case msg := <-client.C:
	//	fmt.Println("开始写")
	//	conn.Write([]byte(msg))
	//}
}

//给

//处理map中的消息 当收到消息后就 给map中的每个都发送消息
func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息的话 这里会堵塞
		fmt.Println("接收到了信息：", msg)
		//给每个客户端发送消息
		for _, value := range onlineMap {
			//给map中 vaLue中的管道发送消息
			value.C <- msg
		}
	}
	defer fmt.Println("完了")

}
