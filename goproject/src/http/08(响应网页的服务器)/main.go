package main

import (
	"net/http"
	"fmt"
)

func myHandler(h http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(h, "hello world")
}

func main() {
	http.HandleFunc("/go", myHandler)
	//在指定的地址进行监听，开启一个http
	http.ListenAndServe("192.168.56.1:2222", nil)
}
