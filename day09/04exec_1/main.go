package main
// n 个台阶有多少个不同的方法
// 一次上一个台阶
// 一次上两个台阶

import "fmt"

func jumpfloor(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var x, y int
	x = 1
	y = 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

func main(){
	fmt.Println(jumpfloor(1))
	fmt.Println(jumpfloor(2))
	fmt.Println(jumpfloor(3))
	fmt.Println(jumpfloor(4))
	fmt.Println(jumpfloor(5))
}