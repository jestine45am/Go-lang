package main

import (
	"fmt"
)

func main() {
	var ch0 = make(chan int, 10)
	var ch1 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch0 <- i:
		case ch1 <- i:
		}
	}
	close(ch0)
	close(ch1)
	for j := range ch0 {
		fmt.Println(j)
	}
	for j := range ch1 {
		fmt.Println(j)
	}

}