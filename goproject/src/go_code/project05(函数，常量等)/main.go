package main
import(
	"fmt"
)


var  c, python, java bool

const bd = "这个是常亮"

const(
	const1="111"
	const2="222"
)



func main(){
	fmt.Println( add(1,2))
	fmt.Println(add_(2,3))
	a,b:=returnMany("123","abc")
	fmt.Println(a,b)
	fmt.Println(split(17))
	fmt.Println(c, python, java)

	fmt.Println(bd,const1,const2)

}


func add(x int,y int)  int {
	return x+y
}


// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add_(x,y int) int {
	return x*y
}



//多个值返回
func returnMany(x,y string) (string,string){
	return x,y
}


//  Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//  返回值的名称应当具有一定的意义，它可以作为文档使用。
//  没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
//  直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
//返回的时候 返回 x,y int 类型
func split(sum int) (x,y int)  {
	x = sum*4/9
	y= sum - x
	return
}




