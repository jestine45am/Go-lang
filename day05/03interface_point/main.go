package main

import (
	"fmt"
)
/*
// 结构体变量和结构体指针
type cat struct{
	name string
}
*/
type cat string
func (c *cat) changeName(){
	*c = "Mike"
}
type dog struct{
	name string
}
func (d dog) changeName(){
	d.name = "Hobby"
}
type animal interface{
	changeName()
}

func f(a animal){
	a.changeName()
}
func main(){
	var c1 cat = "Tom"
	var c2 cat = "Jim"
	var c3 = &c2

	var d1 = dog{"John"}
	var d2 = &dog{"Harry"}

	f(&d1)
	f(d2)
	f(&c1)
	f(&c2)
	f(c3)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(*c3)
	fmt.Println(d1.name)
	fmt.Println((*d2).name)
}
/*
type animal interface{
	changeName()
}

func f(a animal){
	a.changeName()
}
*/