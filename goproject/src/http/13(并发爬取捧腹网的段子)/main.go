package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"regexp"
)

//并发爬取 经典笑话 的段子
// http://www.xiao688.com/cms/list/typeid-2_page-1.html
func main() {

	//段子的标题：<div class="title"><h3><a href=  标题  </a></h3></div>

	//段子的内容： <div class="content" style="margin-top:15px"><div class="content_ads mt0"><script type="text/javascript">ads("joke_content");</script></div><div class="ads_jc_r"></div> 内容 </div>

	var start, end int

	fmt.Printf("请输入起始页(>=1)")

	fmt.Scan(&start)

	fmt.Printf("请输入结束页(>=起始页)")

	fmt.Scan(&end)

	work(start, end)

}

func work(start, end int) {

	page := make(chan int)

	//开始爬取网页
	for i := start; i <= end; i++ {
		go goStartNet(i, page)
	}

	for {
		select {
		case index := <-page:
			fmt.Printf("第%d页爬取完成\n", index)
			break
		}
	}

}

func goStartNet(i int, page chan<- int) {
	url := "http://www.xiao688.com/cms/list/typeid-2_page-" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页，地址是：%s\n", i, url)
	//爬取页面内容
	resp, error := HttpGet(url)
	if error != nil {
		fmt.Println(error)
		return
	}

	//将结果写入到文件中去
	fileName := strconv.Itoa(i) + ".html"
	_, e := os.Create("./goproject/src/http/13(并发爬取捧腹网的段子)/" + fileName)
	if e != nil {
		fmt.Println(e)
		return
	}

	//将爬取的数据进行筛选得到想要的信息
	// <div class="title"><h3><a href=  标题  </a></h3></div>
	//?s遇到\n也会处理的
	re1 := regexp.MustCompile(`<div class="title"><h3><a href=(.*)</a></h3></div>`)

	//re1 := regexp.MustCompile(`<div class="title"><h3><a href=(?s:(.*?))</a></h3></div>`)

	if re1 == nil {
		fmt.Println("匹配错误了")
		return
	}
	content := re1.FindAllStringSubmatch(resp, -1)

	for _, v := range content {
		re2 := regexp.MustCompile(`"(.*)" title=`)
		suContent := re2.FindAllStringSubmatch(v[1], -1)
		//根据url爬取具体的代码
		title, con, errs := spiderOneJoy(suContent[0][1])
		if errs != nil {
			fmt.Println(errs)
			continue
		}
		fmt.Println(title, con)
	}

	page <- i

}

//开始爬取
func HttpGet(url string) (res string, er error) {
	resp, error := http.Get(url)
	defer resp.Body.Close()
	if (error != nil) {
		er = error
		return
	}
	defer resp.Body.Close()

	//读取网页的内容

	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		res += string(buf[:n])
	}

	return
}

//根据返回里面的url 返回具体的内容
func spiderOneJoy(onePath string) (title, con string, err error) {

	resp, error := http.Get("http://www.xiao688.com" + onePath)
	fmt.Println("http://www.xiao688.com" + onePath)
	defer resp.Body.Close()
	if error != nil {
		err = error
		return
	}

	buf := make([]byte, 1024*4)

	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		con += string(buf[:n])
	}
	title = "宋佳宾"

	return

}
