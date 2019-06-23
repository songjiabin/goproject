package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(strings.Contains("chain", "ch"))
	fmt.Println(strings.Contains("ch", "chain"))
	fmt.Println(strings.Contains("chain", ""))  //true
	fmt.Println(strings.Contains("", ""))       //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我")) //true
	fmt.Println("---------------")
	containsAny()
	fmt.Println("---------------")
	Count()
	fmt.Println("---------------")
	index()
}

//func ContainsAny(s, chars string) bool这个是查询字符串中是否包含多个字符.
func containsAny() {
	fmt.Println(strings.ContainsAny("chain", "b"))     // false
	fmt.Println(strings.ContainsAny("chain", "c & i")) // true
	fmt.Println(strings.ContainsAny("chain", ""))      // false
	fmt.Println(strings.ContainsAny("", ""))           // false
}

// 计算里面的数量
func Count() {
	fmt.Println(strings.Count("chainn", "nn")) //1
	fmt.Println(strings.Count("chainn", "n"))  //2
	fmt.Println(strings.Count("chain", ""))    // before & after each rune result:6
}

// Index

func index() {
	fmt.Println(strings.Index("chain", "h"))  //1
	fmt.Println(strings.Index("chainn", "n")) //4
	fmt.Println(strings.Index("chainn", "q")) //-1
	fmt.Println(strings.Index("我是中国人", "中"))  // 返回 6

	s := "我是中国人"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))
	fmt.Println(s[0:15])
	nameRune := []rune(s)
	fmt.Println(nameRune)
	fmt.Println(len(nameRune))
	fmt.Println("string = ", string(nameRune[0:5]))
	fmt.Println(strings.Index(string(nameRune), "中"))

	fmt.Println(strings.IndexByte("hello chain", 'c')) //6
	fmt.Println(strings.IndexByte("hello chain", 'b')) //-1

	fmt.Println(strings.IndexFunc("achainaa", split)) //2
	fmt.Println(strings.IndexFunc("chbin", split))    //-1



	mapDemoth()
}

func split(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}

func mapDemoth() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
