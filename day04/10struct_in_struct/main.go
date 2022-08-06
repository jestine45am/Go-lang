package main

import "fmt"
/*

type address struct{
	provice string
	city string
}
type student struct{
	name string
	age int
	addr address
}

func main(){
	var s1 = student{
		"Tom",
		18,
		address{
			"jiangxi",
			"ganzhou",
		},
	}
	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.addr.provice)
}
*/
// 匿名嵌套结构体
type address struct{
	provice string
	city string
}

type student struct{
	name string
	age int
	address
}

func main(){
	var s1 = student{
		"Tom",
		18,
		address{
			"jiangxi",
			"ganzhou",
		},
	}

	fmt.Println(s1.name)
	fmt.Println(s1.city)
	fmt.Println(s1.address.city)
}