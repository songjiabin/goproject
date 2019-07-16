package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	str := `{"company":"北京邮电大学","subjects":["名人","卡卡西"],"isok":true,"price":90.9}`

	//需要解析到的map中
	m := make(map[string]interface{}, 4)

	error := json.Unmarshal([]byte(str), &m)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(m)

	s := m["isok"]
	su:=m["subjects"]
	fmt.Println(su)
	fmt.Println(s)

}
