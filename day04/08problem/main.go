package main

import "fmt"

type human struct{
	name string
	age int
	gender string
}
func newHuman(n string, a int, g string) human{
	/*
	return human{
		name: n,
		age: a,
		gender: g,
	}
	*/
	return human{n,a,g}
}
func main(){
	var h1 = human{"Tom",18,"male"}
	var h2 = newHuman("Jim",19,"male")
	fmt.Println(h1)
	fmt.Println(h2)
	var h3 human
	h3.name = "John"
	h3.age = 20
	h3.gender = "male"
	fmt.Println(h3)
	var h4 = human{
		name: "Kate",
		age: 17,
		gender: "female",
	}
	fmt.Println(h4)

}