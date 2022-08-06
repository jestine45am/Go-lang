package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func hello(i int){
	defer wg.Done()
	fmt.Println("hello world!", i)
}
func main(){
	/*
	for i := 0; i < 1000; i++ {
		// go hello(i)
		func(i int){
			fmt.Println(i)
		}(i)
	}
	*/
	for i := 0; i < 10; i++{
		wg.Add(1)
		go hello(i)

	}
	fmt.Println("complete!")
	wg.Wait()
}