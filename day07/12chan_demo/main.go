package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){
	var ch0 = make(chan int, 10)
	var ch1 = make(chan int, 10)
	wg.Add(2)
	go func(){
		wg.Done()
		for i := 0; i < 10; i++{
			ch0 <- i
		}
		close(ch0)
	}()
	go func(){
		wg.Done()
		for {
			i,ok := <- ch0
			if !ok {
				break
			} else {
				ch1 <- i
			}
		}
		close(ch1)
	}()
	for j :=  range ch1{
		fmt.Println(j)
	}
	wg.Wait()
}