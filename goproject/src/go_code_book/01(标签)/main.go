package main

import "fmt"

func main() {
ABC:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 2 {
				break ABC
			}
		}
	}

	fmt.Println("-------")



	b:=[]byte("are you ok ")
	fmt.Println(b)


}
