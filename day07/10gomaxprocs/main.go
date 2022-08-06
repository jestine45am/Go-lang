package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a(){
	defer wg.Done()
	for i := 0; i < 10; i++{
		fmt.Printf("A:%d\n",i)
	}
}

func b(){
	defer wg.Done()
	for i := 0; i < 10; i++{
		fmt.Printf("B:%d\n",i)
	}
}

func main(){
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	fmt.Println(n)
	wg.Add(2)
	a()
	b()
	wg.Wait()
}