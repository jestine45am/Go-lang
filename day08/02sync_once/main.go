package main

import (
	"fmt"
	"sync"
)
var (
	wg sync.WaitGroup
	lock sync.Mutex
	so sync.Once
	count int
)
func f(x chan int){
	defer wg.Done()
	for i := 0; i < 30; i++ {
		x <- i
	}
	lock.Lock()
	count++
	lock.Unlock()
	if count == 3 {
		so.Do(func(){close(x)})
	}
}
func main(){
	ch := make(chan int, 100)
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go f(ch)
	}
	wg.Wait()
	for v := range ch {
		fmt.Println(v)
	}
}