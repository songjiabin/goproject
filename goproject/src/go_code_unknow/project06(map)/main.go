package main

import (
	"fmt"
	"sort"
)

//map
func main() {

	//1、
	var m map[string]string
	fmt.Println(m)

	//2、
	m1 := map[int]string{}
	fmt.Println(m1)

	//3、
	var m2 map[int]string
	m2 = make(map[int]string)
	fmt.Println(m2)

	//4、
	m3 := make(map[int]string)
	fmt.Println(m3)

	m4 := make(map[int]map[int]string)
	m4[0] = make(map[int]string)
	m4[0][2] = "1"
	fmt.Println(m4)

	//多返回值
	// a取出的数据 ok是否有
	a, ok := m4[2][1]
	if !ok {
		m4[2] = make(map[int]string)
		m4[2][1] = "2"
	}
	fmt.Println(a, ok)
	fmt.Println(m4)

	fmt.Println("-----------------")

	//map 的遍历
	sm := make([]map[int]string, 5)
	fmt.Println(sm)

	for i, _ := range sm {
		sm[i] = make(map[int]string)
		sm[i][0] = "ok"
	}
	fmt.Println(sm)

	sm2 := make(map[int]string)
	sm2[0] = "1"
	sm2[1] = "2"
	for k, v := range sm2 {
		fmt.Println(k, v)
	}

	//对map key进行排序

	ak := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	aa := make([]int, len(ak))

	i := 0
	for key, _ := range ak {
		aa[i] = key
		i++
	}
	//对数组aa进行排序
	fmt.Println(aa)
	sort.Ints(aa)
	fmt.Println(aa)

	//然后根据key有序的取出map的值

	for i, v := range aa {
		fmt.Println(i, ak[v])
	}

	//将map1中的 key value 进行倒换
	fmt.Println("----------------------------")
	map1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(map1)

	map2 := make(map[string]int)

	for k, v := range map1 {
		map2[v]=k
	}

	fmt.Println(map2)

}
