package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	fileName := strconv.Itoa(i) + ".txt"
	file, e := os.Create(fileName)
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

		//将数据写入到文件中去
		file.WriteString(title + "\n" + con)
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

	t := regexp.MustCompile(`</a></div><h1>(.*)</h1></div>`)
	content := t.FindAllStringSubmatch(con, -1)

	//fmt.Println(content)

	c := regexp.MustCompile(`</script></div><div class="ads_jc_r"></div>(.*)</div>`)
	con_con := c.FindAllStringSubmatch(con, -1)

	con = con_con[0][1]
	title = content[0][1]

	con = strings.Replace(con, "\t", "", -1)
	con = strings.Replace(con, "\n", "", -1)
	con = strings.Replace(con, "\r", "", -1)
	con = strings.Replace(con, "<br />", "", -1)

	title = strings.Replace(title, "\t", "", -1)
	title = strings.Replace(title, "\n", "", -1)
	title = strings.Replace(title, "\r", "", -1)
	title = strings.Replace(title, "<br />", "", -1)
	return




}




