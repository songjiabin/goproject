package main

import "fmt"

func noSame(data []string) []string {
	out := data[:1]
	// 遍历原始切片字符串
	for _, word := range data {

		// 比较取出的 word 是否在 out 中存在 -- for
		for i := 0; i < len(out); i++ {
			if word == out[i] {
				break
			}
		}

		out = append(out, word)

	}
	return out
}

func main() {
	data := []string{"red", "black", "red", "yellow", "yellow", "pink", "blue", "pink", "blue"}

	afterData := noSame(data)
	fmt.Println("Afterdata:", afterData)

}
