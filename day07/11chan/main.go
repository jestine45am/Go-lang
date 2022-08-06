package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main(){
	a := make(chan int)
	wg.Add(1)
	go func(){
		defer wg.Done()
		x := <- a
		fmt.Println(x)
	}()
	a <- 10
	fmt.Println(a)
	wg.Wait()
	// a = make(chan int,10)  // 0xc0000200b0
	// fmt.Printf("%v\n", a)
}