package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

func main() {

	// https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0 翻页就是 +50

	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50

	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100

	//起始页 ，结束页
	var start, end int

	fmt.Printf("请输入起始页(>=1)")

	fmt.Scan(&start)

	fmt.Printf("请输入结束页(>=起始页)")

	fmt.Scan(&end)

	work(start, end)

}

func work(start, end int) {
	if end < start {
		fmt.Println("起始页不能等于结束页")
		return
	}

	//开始爬取网页
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa(i*50-50)
		content, error := SpiderPage(url)
		if error != nil {
			fmt.Println(error)
			continue
		}
		//将内容写入到本地的一个文件
		fileName := "./goproject/src/http/11(百度贴吧爬虫)/" + strconv.Itoa(i) + ".html"
		f, errorFile := os.Create(fileName)
		if errorFile != nil {
			fmt.Println(errorFile)
			continue
		}
		f.Write([]byte(content))
		f.Close()
	}

}

//开始爬取
func SpiderPage(url string) (result string, err error) {
	resp, errorHttp := http.Get(url)
	if errorHttp != nil {
		err = errorHttp
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		result += string(buf[:n])
	}

	return

}
