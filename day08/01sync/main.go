package main

/*
// 利用共享内存的方式进行数据交互往往容易产生竞态问题
import (
	"sync"
	"fmt"
)
var x int = 0
var wg sync.WaitGroup
func add(){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		x++
	}
}

func main(){
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
*/

/*
// 利用通道来实现数据的共享

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var ch = make(chan int, 100)
func add( ch chan int){
	defer wg.Done()
	for i:=0; i<10000; i++ {
		for {
			select {
			case x := <- ch: {
					x++
					ch <- x
					goto r1
				}
			default: goto r2
			}
			r2:
		}
		r1:
	}
}
func main(){
	ch <- 0
	wg.Add(2)
	go add(ch)
	go add(ch)
	wg.Wait()
	x := <- ch
	fmt.Println(x)
}
*/

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mt sync.Mutex
	x int
)

func add(){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		mt.Lock()
		x++
		mt.Unlock()
	}
}
func main(){
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}