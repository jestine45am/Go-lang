package main

import "fmt"

func main(){
/*
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(*p)
	*/
/*
	var p *int
	fmt.Printf("%v\n",p)
	*p = 100
	fmt.Println(*p)
*/

	// new() 函数开辟内存空间，返回一个指针
	var p = new(int)
	*p = 100
	fmt.Println(*p)
}