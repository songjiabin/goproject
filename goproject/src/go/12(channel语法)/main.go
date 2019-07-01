package main

func main() {
	//demo1()
	//demo2()
	demo3()
}


//单方向的只能读
func demo2() {
	ch := make(chan int)
	var chRead <-chan int = ch

	<-chRead

}

func demo1() {
	//双方向的 channel
	ch := make(chan int)
	//双方向的channel能隐士转换为单方向的channel
	//单方向的只能写
	var chWrite chan<- int = ch
	chWrite <- 666
}

func demo3()  {

}