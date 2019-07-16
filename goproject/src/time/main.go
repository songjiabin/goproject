package main

import (
	"time"
	"fmt"
	"err"
)

func main() {

	//获取当前的时间
	t := time.Now()
	fmt.Println(t) //2019-07-16 17:12:38.9511934 +0800 CST m=+0.003924701

	//手动获取指定的时间
	datenow := time.Date(2019, 12, 12, 11, 12, 34, 0, time.Local)
	fmt.Println(datenow)

	//格式化时间
	//年月日 时分秒
	s1 := t.Format("2006年1月2日 15:04:05") //注意规定是这个时间
	fmt.Println(s1)

	//年月日
	s2 := t.Format("2006/1/2")
	fmt.Println(s2)

	//解析为时间类型 string->time
	t3 := "2019年10月11日" //string
	parse, e := time.Parse("2006年1月2日", t3)
	err.PrintErr(e)
	fmt.Println(parse)
	fmt.Printf("类型是：%T\n", parse)

	//获取年月日
	fmt.Println(t.String())
	year, month, day := t.Date()
	fmt.Println(year, month, day)

	//获取时分秒
	hour, min, sec := t.Clock()
	fmt.Println(hour, min, sec)

	fmt.Println(t.YearDay()) //过了多少天了
	fmt.Println(t.Weekday()) //周几呢

	//获取时间戳 从1970年。。。到现在的秒数
	fmt.Println(t.Unix())

	//获取纳秒 从1970年。。。到现在的纳秒数
	fmt.Println(t.UnixNano())

	//时间间隔
	fmt.Println(t.Add(time.Minute)) //比你设定的时间 多一分钟

	fmt.Println(t.AddDate(1, 0, 0))//比设定的时间多一年


	//计算时间 当前的时间 - 当前的时间+1分钟
	fmt.Println(t.Sub(t.Add(time.Minute))) //结果是；-1m0s 一分钟


}
