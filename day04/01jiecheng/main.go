package main

import "fmt"

func f(n int64)(m int64) {
	if n > 1 {
		m = n * f(n-1)
	}else if n == 1{
		m = 1
	}else {
		m = 0
	}
	return m
}
func main(){
	x := f(0)
	fmt.Println(x)
}