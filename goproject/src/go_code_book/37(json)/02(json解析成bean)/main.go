package main

import (
	"encoding/json"
	"fmt"
)

type JsonBean struct {
	Company  string   `json:"company"` //给json修改名字
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	//解析json 冰封装成bean
	str := `{"company":"北京邮电大学","subjects":["名人","卡卡西"],"isok":true,"price":90.9}`
	var jsonBean JsonBean
	error := json.Unmarshal([]byte(str), &jsonBean)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(jsonBean)

	//如果只需要里面的一个参数
	type oneParams struct {
		Subjects []string
	}

	var oneparams oneParams
	error_ := json.Unmarshal([]byte(str), &oneparams)
	if error_ != nil {
		fmt.Println(error_)
		return
	}
	fmt.Println(oneparams)

}
