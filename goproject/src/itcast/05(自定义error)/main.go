package main

import (
	"fmt"
	"time"
)

func main() {
	utc := time.Now().UTC()
	fmt.Print(utc)
}
