package main

import "fmt"

func main() {

	header:="Bearer 宋佳宾"
	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)

	fmt.Println(t)
}
