package main

import (
	"fmt"
	"path"
)

//常量：
const name bool = true

const name2 = 113232

func main() {

	fmt.Println(name)
	fmt.Println(name2)

	x := 6 + 2i
	fmt.Println(x * x)

	var x2 [10]int
	x2[1] = 23
	fmt.Println(x2)

	x3 := [2]int{2, 4}
	fmt.Println(x3)



	db:="/c:/abc"
	fmt.Println(path.Ext(db))


}
