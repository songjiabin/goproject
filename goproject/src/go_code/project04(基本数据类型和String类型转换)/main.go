package main

import(
	"fmt"
	"strconv"
	"math/rand"
	"math"
)

//基本数量类型和String的转换
// fmt 包下面的%t/d/f/c 通過   https://studygolang.com/pkgdoc
func main()  {
	//方式一：
	// method_1()
	//方式二：
	method_2()



    fmt.Println("My ...",rand.Intn(10))
	fmt.Println("My ...",math.Pi)

}


func method_1()  {
	var num1 int = 80
	var num2 float64 = 0.254
	var b bool = true
	var myChar byte = 'h'
	
	var str string

	//int => String的转换
	str =fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q \n",str,str)


	// float => String
	str =fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str= %q \n",str,str)


	// bool => string
	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q \n", str,str)


	// char => string 
	str=fmt.Sprintf("%c",myChar)
	fmt.Printf("str stye %T str=%q \n",str,str)
}


func method_2()  {
	var num1 int = 80
	var num2 float64 = 0.254
	var b bool = true
	// var myChar byte = 'h'
	var str string


	// int => string
	str=strconv.FormatInt(int64(num1),10)
	fmt.Printf("str type %T str=%q \n",str,str)

	// float => string
	// f=>显示的格式  
	// 10=>表示小数点保留10位
	// 64=>表示小数64
	str =strconv.FormatFloat(num2,'f', 10,64)
	fmt.Printf("str stye %T str=%q \n",str,str)

    //bool => string
    str =strconv.FormatBool(b)
	fmt.Printf("str stye %T str=%q \n",str,str)


	// int = > string
	// Itoa是FormatInt(i, 10) 的简写。
	str =strconv.Itoa(num1)
	fmt.Printf("str style %T st=%q \n",str,str)

}