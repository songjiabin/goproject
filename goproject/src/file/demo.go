package main

import (
	"os"
	"fmt"
)

func main() {
	remove := os.Remove("./abc.jpgtemp.txt")
	fmt.Println(remove)
}
