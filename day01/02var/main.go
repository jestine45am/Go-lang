package main

import "fmt"

// 声明变量
var name string
var age uint
var isOk bool

// 批量声明变量
/*
var (
	name1 string
	age1 uint
	isOk1 bool
)
// 变量声明之后，未赋值时，那么其值时对应变量的空值
*/
// 变量必须先声明再使用
// 变量声明之后必须使用，否则会编译不了
// 对于全局变量来说，可能在别的包中会使用到，所以，即便在本包中未使用也可以完成编译
// 但是对应局部变量来说，必须马上使用，否则是无法完成编译


func main() {
	name = "lilei"
	age = 18
	isOk = true

	fmt.Print(isOk)
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

}
