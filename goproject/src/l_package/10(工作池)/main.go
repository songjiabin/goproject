package main

import (
	"time"
	"sync"
	"math/rand"
	"fmt"
)

//工作池
//工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
//我们会使用缓冲信道来实现工作池。我们工作池的任务是计算所输入数字的每一位的和。例如，如果输入 234，结果会是 9（即 2 + 3 + 4）。向工作池输入的是一列伪随机数。

var jobs chan Job
var results chan Result

func main() {
	jobs = make(chan Job,10)
	results = make(chan Result,10)

	startTime := time.Now()


	//一共有100个数值要进行计算
	noOfJobs := 100
	//创建数据并将数据交给通道中
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)


	//处理的线程池子
	noOfWorkers := 2
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

type Job struct {
	id int
	//需要计算的
	randomno int
}

type Result struct {
	job Job
	//需要计算的结果
	sumofdigits int
}

//将作业分配给工作者
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		//每次只能写10个，因为容量最大就是10
		jobs <- job
	}
	fmt.Println("完成要close()了-----------------------------")
	close(jobs)
}

//创建一个工作池
// noOfWorkers 工作协程的数量
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	//等待所有的 Go 协程执行完毕。所有协程完成执行之后，函数会关闭 results 信道
	wg.Wait()
	close(results)
}

//工作线程
func worker(wg *sync.WaitGroup) {
	//实时循环jobs线程，一旦有工作
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

//计算
func digits(num int) int {
	sum := 0
	no := num
	for no != 0 {
		//计算位数的和
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func result(done chan bool) {
	//如果没有数据就一直堵塞这
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
