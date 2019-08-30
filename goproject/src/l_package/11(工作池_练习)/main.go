package main

import (
	"math/rand"
	"fmt"
)

//计算出每个数字上每位上的和
//123 ->1+2+3

type Job struct {
	id     int
	number int
}

type Result struct {
	job Job
	sum int
}

var jobs chan Job
var results chan Result
var isOk chan bool

func main() {
	isOk = make(chan bool)
	jobs = make(chan Job, 10)
	results = make(chan Result, 10)

	nums := 10
	go getNums(nums)
	go getResult()

	//创建pool
	handleJob()

	<-isOk
}

//取Job 并处理
func handleJob() {
	//监听取出
	for job := range jobs {
		//取出里面的job
		result := Result{job: job, sum: digits(job.number)}
		results <- result
	}
	fmt.Println("结束,results")
	close(results)

}

//生成数值 并写入通道中
func getNums(nums int) {
	for i := 0; i < nums; i++ {
		rondom := rand.Intn(999)
		jmodel := Job{id: i, number: rondom}
		jobs <- jmodel
	}
	//当读完以后就可以关闭了
	fmt.Println("结束,jobs")
	close(jobs)
}

func getResult() {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.number, result.sum)
	}
	isOk <- true
}

//将每个位相加法  并得到结果
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	return sum
}
