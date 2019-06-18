package main
import (
	"fmt"
	"unsafe"
)


//变量三要素 ： 变量名字 + 值 + 类型  



//全局定义变量
//主要：中间没有逗号呢
var (
	n7="全局变量哦"
	n8=100
	n9=0.18
)


func main(){
	 
   //指定的变量的类型
   var i int =1;
   fmt.Println("i=",i)


   //会根据你的值自行推断出类型
   var i2 = 00.94
   fmt.Println("i=",i2)


   //声明并且赋值 
   //相当于 var name String ="宋佳宾"
	name:="宋佳宾"
	fmt.Println(name)



	//多变量的声明(相同的类型)
	var n1,n2,n3 int ;
	fmt.Println(n1,n2,n3)


	//多变量的声明（不同的类型）
	var n4,n5,n6 = 1,"宋佳宾",0.25;
	fmt.Println(n4,n5,n6)



	//全局变量的声明
	fmt.Println(n7,n8,n9)



   //fmt.printf 用于格式化的输出
   // unsafe 查看字节数 
   var ns int64= 10
   fmt.Printf("ns的类型是：%T ns占用的字节数是：%d \n",ns,unsafe.Sizeof(ns))


  //使用：声明

  b:="宋佳宾"
  fmt.Println(b)


}