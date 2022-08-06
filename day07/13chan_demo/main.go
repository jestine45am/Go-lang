package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func work(n int, jobs chan<- int){
	defer wg.Done()
	for i := 1; i <= n; i++{
		jobs <- i
	}
	close(jobs)
}
func proc(id int, jobs <-chan int, results chan<- int){
	defer wg.Done()
	for {
		j,ok:= <- jobs
		if !ok {
			break
		}
		fmt.Printf("worker id:%d start %d\n", id, j)
		fmt.Printf("worker id:%d end %d\n", id, j)
		results <- j * 2
	}
}

func main(){
	var jobs = make(chan int, 10)
	var results = make(chan int, 10)
	wg.Add(4)
	go work(10, jobs)
	for w := 1; w <= 3; w++ {
		go proc(w, jobs, results)
	}

	for a := 1; a <= 10; a++{
		<-results
	}
	wg.Wait()

}