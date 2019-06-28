package _1_生成json_

import (
	"encoding/json"
	"fmt"
)

/**
{
    "Company":"北京邮电大学",
    "Subjects":[
        "名人",
        "卡卡西"
    ],
    "Isok":true,
    "Price":90.9
}
 */

type JsonBean struct {
	Company  string   `json:"宋佳宾"`     //给json修改名字
	Subjects []string `json:"-"`       //如果后面是：- 表示不被输出
	Isok     bool     `json:",string"` //将此格式装换成字符串进行输出
	Price    float64
}

func main() {

	createJson1()
	createJson2()

}
func createJson1() {
	s := JsonBean{
		Company: "北京邮电大学",
		Subjects: []string{
			"名人", "卡卡西",
		},
		Isok:  true,
		Price: 90.9,
	}
	//a, b := json.Marshal(s)
	a, b := json.MarshalIndent(s, "", "	")
	//格式化编码的格式
	fmt.Println(string(a), b)
}

func createJson2() {
	m := make(map[string]interface{}, 4)
	m["Company"] = "北京邮电大学"
	m["Subjects"] = []string{"名人", "卡卡西"}
	m["Isok"] = true
	m["Price"] = 90.3

	result, error := json.MarshalIndent(m, "", " ")
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(string(result))
}
