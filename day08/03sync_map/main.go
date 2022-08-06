package main

import (
	"fmt"
	"sync"
)
var (
	a1 = []string{"one","two","three","four","five","six","seven","eight","nine","ten"}
	b1 = []int {1,2,3,4,5,6,7,8,9,10}
	wg sync.WaitGroup
	m = sync.Map{}
) 
func set(){
	defer wg.Done()
	for i := 0; i < 10; i++{
		m.Store(a1[i], b1[i])
	}
}
func get(){
	defer wg.Done()
	for i := 0; i < 10; i++{
		value,ok := m.Load(a1[i])
		if ok {
			fmt.Printf("m[%s]:%d\n", a1[i], value)
		}
	}
}
func main(){
	for i := 0; i < 200; i++{
		wg.Add(1)
		go set()
	}
	for j := 0; j < 200; j++{
		wg.Add(1)
		go get()
	}
	wg.Wait()
}