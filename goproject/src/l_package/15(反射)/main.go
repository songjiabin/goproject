package main

//反射

import (
	"fmt"
	"reflect"
)

type order struct {
	name       string
	ordId      int
	customerId int
}

func main() {
	o := order{ordId: 100, customerId: 1, name: "宋佳宾"}
	getType(o)
}

func getType(params interface{}) {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	kind := t.Kind()
	fmt.Println("type", t)
	fmt.Println("value", v)
	fmt.Println("kind", kind) //表示特定的实际类型

	if reflect.ValueOf(params).Kind() == reflect.Struct {

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
