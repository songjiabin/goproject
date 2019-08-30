package main

import "fmt"

func main() {
	f := func() int {
		return 3
	}

	fmt.Println(f())
}
