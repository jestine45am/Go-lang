package main
import "fmt"

// 匿名函数定义
func main(){
	/*
	f1 := func(x,y int){
		fmt.Println(x+y)
	}
	f1(1,10)
	*/
	func(x,y int){
		fmt.Println(x+y)
	}(100,200)
}