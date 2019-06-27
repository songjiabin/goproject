package main

func testA() {
	//程序直接崩溃
	panic("this is a panic") //提示崩溃的原因
}

func testB() {

}

//显示调用的话 会导致程序中断
func main() {
	testA()
}
