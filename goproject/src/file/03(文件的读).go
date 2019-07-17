package main

import (
	"os"
	"fmt"
	"io"
)

func main() {

	filePath := "./a.txt"

	file, e := os.Open(filePath)
	if e != nil {
		fmt.Println(e)
		return
	}
	buf := make([]byte, 4, 4)

	//第一次读
	n, err := file.Read(buf)
	fmt.Println(n, err, string(buf))

	//第二次读
	n, err = file.Read(buf)
	fmt.Println(n, err, string(buf))

	//第三次读取
	n, err = file.Read(buf)
	fmt.Println(n, err, string(buf))

	//第四次读
	n, err = file.Read(buf)
	fmt.Println(n, err, string(buf))

	//读取的数据是 abcdefghi
	//分四次读取 当最后一次读取的时候 err为 EOF

	fmt.Println("-----------------------------------")
	meth2()
}

func meth2() {
	//所以采用循环的读法
	filePath := "./a.txt"
	file, e := os.Open(filePath)
	if e != nil {
		fmt.Println(e)
		return
	}
	bufs := make([]byte, 4, 4)
	for {
		ns, errs := file.Read(bufs)
		if ns == 0 || errs == io.EOF {
			fmt.Println("读取到了末尾")
			break
		}
		fmt.Println(string(bufs[:ns]))
	}
}
