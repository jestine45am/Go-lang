package main

import "fmt"

type myInt int
func (m myInt) f1(){
	fmt.Println("I am a int type")
}
func main(){
	m := myInt(100)
	m.f1()
}