package main

import "fmt"

func main() {
	str := []string{"red", "black", "red", "yellow", "yellow", "pink", "blue", "pink", "blue"}

	startClie := str[:1]

	for _, v := range str {
		i := 0
		for ; i < len(startClie); i++ {
			if startClie[i] == v {
				break
			}
		}

		//当循环完成 再相加
		if i == len(startClie) {
			startClie = append(startClie, v)
		}


	}

	fmt.Println(startClie)

	strings := make(map[int]string, 5)
	fmt.Println(len(strings))
}


