package main

import "fmt"

type human struct{
	name string
	age int
	gender string
}
/*
func f_struct(a int, b,c string)(human){
	return struct{
		name string;
		age int;
		gender string
	}{b,a,c}
}
*/
/*
// 通过结构体传递
func f_struct1(a int, b,c string)human{
	return human{
		name: b,
		age: a,
		gender: c,
	}
}
func main(){
	h1 := f_struct1(18,"Tom","male")
	fmt.Println(h1)

}
*/
// 通过结构体指针传递
func f_struct2(a int, b, c string)*human{
	return &human{
		name: b,
		age: a,
		gender: c,
	}
}
func main(){
	h2 := f_struct2(29,"John","male")
	fmt.Println(*h2)
}