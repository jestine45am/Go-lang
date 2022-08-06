package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
// job 结构体，用于保存随机数
type job struct{
	id int64
}
// result 结构体，用于保存随机数和各个位数的和
type result struct{
	job 
	sum int64
}
var wg sync.WaitGroup
// 开启一个 goroutine 循环生成 int64 随机数，并将结果发送到 jobchan
func RandProduce(jobchan chan<- *job) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixMicro())
		x :=  rand.Int63()
		newJob := &job{x} 
		jobchan <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}
// 开启一个 goroutine 从 jobchan 中取值，并计算各个位数的和，最后将结果发送到 resultchan
func RandCal(jobchan <-chan *job, resultchan chan<- *result){
	defer wg.Done()
	for {
		i,ok := <- jobchan
		if !ok{
			break
		}
		var sum int64 = 0
		r := i.id
		for r > 0 {
			k := r % 10
			r /= 10
			sum += k
		}
		newResult := &result{*i,sum}
		resultchan <- newResult
	}
	
}

func main(){
	jobchan := make(chan *job, 100)
	resultchan := make(chan *result, 100)
	// 从 resultchan 中取值，并将结果打印
	
	wg.Add(25)
	go RandProduce(jobchan)
	for i := 0; i < 24; i++{
		go RandCal(jobchan,resultchan)
	}
	for result := range resultchan{
		fmt.Println(result.job.id, result.sum)
	}
	wg.Wait()
}