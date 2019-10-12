package config

import (
	"fmt"
	"os"
)

func GetAbc() {

	file, e := os.Open("./conf/config.yaml")
	if e != nil {
		fmt.Print(e)
		return
	}

	buf := make([]byte, 1024*2)
	// n: 读取的长度
	n, error := file.Read(buf);
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(string((buf[:n])))

}
