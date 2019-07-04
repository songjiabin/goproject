package main

import (
	"net/http"
	"fmt"
)

func main() {

	res, error := http.Get("http://192.168.56.1:3333/go.html")
	if error != nil {
		fmt.Println("http get error", error)
		return
	}

	defer res.Body.Close()

	//读取里面的body
	buf := make([]byte, 1024)

	var content string

	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[:n])
	}

	fmt.Println(content)



}
