package main

import "fmt"

type human struct{
	name string
	age int
	gender string
}
/*
type dog struct{
	name string
}
*/

func newHuman(n string,a int, g string)human{
	return human{
		name: n,
		age: a,
		gender: g,
	}
}
/*
func newDog(n string)dog{
	return dog{
		name: n,
	}
}
*/

func (h human) f1(){
	h.age++
	fmt.Println(h.age)
}
func main(){
	h1 := newHuman("John", 20, "male")
	h1.f1()
	fmt.Println(h1.age)
}