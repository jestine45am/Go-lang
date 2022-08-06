package main

import (
	"fmt"
)

type cat struct{
	name string
	feet uint
}
func (c cat)showname(){
	fmt.Println(c.name)
}
func (c cat)showfeet(){
	fmt.Println(c.feet)
}

type animal interface{
	animal01
	animal02
}
type animal01 interface{
	showname()
}
type animal02 interface{
	showfeet()
}
func f0(a animal){
	a.showname()
}
func f1(a animal01){
	a.showname()
}
func f2(a animal02){
	a.showfeet()
}
func main(){
	var c1 = cat{"Tom",4}
	f1(c1)
	f2(c1)
	f0(c1)

}