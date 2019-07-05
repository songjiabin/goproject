package main

import (
	"regexp"
	"fmt"
)

func main() {

	str := `"/cms/article/id-142236.html" title="经典笑话七则" target="_blank">经典笑话七则`


	re2 := regexp.MustCompile(`"(.*)" title=`)
	suContent := re2.FindAllStringSubmatch(str, -1)


	fmt.Println(suContent[0])
	fmt.Println(suContent[0][0])
}
