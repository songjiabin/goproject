package main
import (
	"fmt"
)

// defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func main()  {
	
	//修饰的推迟调用 但是里面的方法会首先调用 大事结果先不打印
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer fmt.Println(add(1,2))

	fmt.Println(add(2,3))


	defer fmt.Println(add(4,5))

	//被推迟的函数会按照后进先出的顺序调用。

   for i := 0; i < 10; i++ {
	   defer fmt.Println(i)
   }


   fmt.Println("最后的....")



}


func add(x,y int) int  {
	fmt.Println(x,y)
	return x+y
}