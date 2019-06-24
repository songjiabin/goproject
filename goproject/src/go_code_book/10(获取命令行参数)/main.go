package main

import "fmt"
import "os"

func main() {

	list := os.Args
	fmt.Println("len=:", len(list))
	//for i := 0; i < len(list); i++ {
	//	fmt.Println(list[i])
	//}

	for _, v := range list {
		fmt.Println(v)
	}





}
