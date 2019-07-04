package main

import "net/http"

//w 给客户端回复数据
//r 读取客户端的数据
func handConn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第一个,hello world,哈哈哈"))
 }

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/go.html", handConn)

	//监听绑定
	http.ListenAndServe("192.168.56.1:3333", nil)
}
