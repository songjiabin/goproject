package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	go func() {
		resp, err := http.Get("http://127.0.0.1:1255")
		if err != nil {
			fmt.Println(resp.StatusCode)
			return
		}
	}()

	time.Sleep(time.Second*10)
}
