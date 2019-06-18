package main

import "fmt"

//定义结构体
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	type1()

	type2()

	type3()

	type4()
}
func type1() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	// 这个是映射  感觉就是键值对  对应啊
	a := make(map[string]string)
	a["a"] = "a"
	fmt.Println(a["a"])
}

//映射的文法
// 映射的文法与结构体相似，不过必须有键名。
func type2() {

	a := map[string]Vertex{
		"1": Vertex{
			Lat:  0.1,
			Long: 0.23},
	}

	fmt.Println(a)

}

// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
func type3() {

	a := map[string]Vertex{
		"1": {Lat: 0.4, Long: 0.1},
		"2": {Lat: 1.1, Long: 2.3},
	}

	fmt.Println(a)

}

//修改映射
func type4() {
	a := make(map[string]int)

	//添加
	a["1"] = 1
	a["2"] = 10
	fmt.Println(a)

	//删除key为1的键值对
	delete(a, "1")
	fmt.Println(a)

	//查看map a中是否存在 key 为2 的值
	// v:如果存在则返回值，不存在则返回该映射的零值
	//ok:bool类型 是否存在该key
	v, ok := a["2"]
	fmt.Println(v, ok)

	b := make(map[int]bool)
	b[1] = true

	v1, ok1 := b[2]
	fmt.Println(v1, ok1)

}
