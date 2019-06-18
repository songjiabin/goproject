package main

import (
	"fmt"
	"strings"
)

//切片
func main() {

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	//切片 意思就是从1到2  包含第一个，不包含最后一个
	b := a[1:2]

	//切片
	var c [] int = a[1:3]

	fmt.Println(b)
	fmt.Println(c)

	type2()
	type3()
	type4()
	type5()
	type6()
	type7()
	type8()
	type9()
}

// 切片就像数组的引用
func type2() {

	fmt.Println("-----------------------------")

	name := [4]string{
		"鸣人",
		"佐助",
		"卡卡西",
		"自立也",
	}

	fmt.Println(name)

	a := name[0:2]
	b := name[1:3]
	fmt.Println(a, b)

	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。
	b[0] = "xxx"
	fmt.Println(a, b)

}

//切片文法
// 切片文法类似于没有长度的数组文法。
func type3() {
	fmt.Println("-----------------------------")

	q := []int{2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}

	fmt.Println(s)

}

//切片的默认行为
func type4() {
	fmt.Println("-----------------------------")
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a = a[1:4]
	fmt.Println(a)

	//默认从头截取两个
	a = a[:2]
	fmt.Println(a)

	//从第一个到最后
	a = a[1:]
	fmt.Println(a)

}

//切片的长度和容量
func type5() {
	// 切片拥有 长度 和 容量。
	//切片的长度就是它所包含的元素个数。
	//切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	fmt.Println("-----------------------------")
	a := []int{1, 2, 3, 4, 5}
	// 截取切片使其长度为 0
	a = a[:0]
	fmt.Println(a)

	/// 拓展其长度
	a = a[:4]
	fmt.Println(a)

	// 舍弃前两个值
	a = a[2:]
	fmt.Println(a)

	//查看 len长度   cap容量
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

}

//切片的 零值  为 nil
func type6() {
	fmt.Println("-----------------------------")
	var s []int
	fmt.Println(s, len(s), cap(s))

	if (s == nil) {
		fmt.Println("nil!")
	}
}

// 用make 创建切片
func type7() {
	/**
		切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
		make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	 */
	fmt.Println("-----------------------------")

	a := make([]int, 5)
	fmt.Println(a)

	//创建数组 容量是5，长度是5
	b := make([]int, 5, 5)
	fmt.Println(b, len(b), cap(b))

	//切片 找到两个
	c := b[:2]
	fmt.Println(c)

	//切片前两个
	d := b[2:5]
	fmt.Println(d)

}

//切片的切片
func type8() {

	fmt.Println("-----------------------------")
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//join的使用
	fmt.Println(strings.Join([]string{"2", "3", "4"}, "_"))

}

func type9() {
	fmt.Println("-----------------------------")
	var s [] int
	fmt.Println(s)

	//添加一个空的切片
	s = append(s, 0)
	fmt.Println(s)


	// 这个切片会按需增长
	s = append(s, 1)
	fmt.Println(s)


	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	fmt.Println(s)

}
