package main

import "fmt"

//类型断言
func main() {

	m := make(map[int]string)

	m[1] = "宋佳宾"

	i := make([]interface{}, 5)

	i = append(i, 1)
	i = append(i, true)
	i = append(i, "true")
	i = append(i, 'c')
	i = append(i, [2]int{1, 3})
	i = append(i, []int{1, 2, 3, 4})
	i = append(i, m)

	for _, v := range i {
		switch v.(type) {
		case int:
			fmt.Println(v, "-->int")
		case bool:
			fmt.Println(v, "-->bool")
		case string:
			fmt.Println(v, "-->string")
		case []int:
			fmt.Println(v, "-->[]int")
		case [2]int:
			fmt.Println(v, "-->[2]int")
		case byte:
			fmt.Println(v,"-->byte")
		case map[int]string:
			fmt.Println(v,"-->map")
		}
	}

}
