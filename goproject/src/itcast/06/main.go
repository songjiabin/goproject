package main

import "fmt"

func main() {

	//m := make(map[int]string)
	//m[1] = "宋佳宾"
	//test(m)
	//fmt.Println(m)

	test2()
}

func test(m map[int]string) {
	m[2] = "许美琳"
	fmt.Println(m)

	delete(m,111)



}

func test2()  {
	strings := make(map[int]string, 2)
	strings[0]="1111"
	strings[2]="2222"
	strings[3]="233"
	fmt.Println(strings)



	v,ok:=strings[0]
	if ok {
		fmt.Println(strings[0])
		fmt.Println(v)
	}


	delete(strings,4)



}