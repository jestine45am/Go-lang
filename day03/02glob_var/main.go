package main

import "fmt"
// 定义全局变量
var x = 100

func f1(){
	x := 10
	fmt.Println(x)
	for x := 5; x < 10; x++{
		fmt.Println(x)
	}
}
func main()  {
	f1()	
}