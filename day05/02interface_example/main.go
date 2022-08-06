package main

import "fmt"

type cat struct{name string}

type dog struct{name string}

type human struct{name string}

func (c cat) move(){
	fmt.Println(c,"is running")
}
func (d dog) move(){
	fmt.Println(d,"is running")
}
func (h human) move(){
	fmt.Println(h,"is running")
}

func (c cat) speak(){
	fmt.Println(c, "is speaking")
}
func (d dog) speak(){
	fmt.Println(d, "is speaking")
}
func (h human) speak(){
	fmt.Println(h, "is speaking")
}

func (h human) talk(){
	fmt.Println(h, "it talking")
}
func (c cat) talk(){
	fmt.Println(c, "it talking")
}
func (d dog) talk(){
	fmt.Println(d, "it talking")
}

type animal interface{
	move()
	speak()
	talk()
}

func f(x animal){
	x.move()
	x.speak()
	x.talk()
}
func main(){
	var c1 = cat{"cat"}
	var d1 = dog{"dog"}
	var h1 = human{"human"}
	var ss animal
	f(ss)
	/*
	f(c1)
	f(d1)
	f(h1)
	*/
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", d1)
	fmt.Printf("%T\n", h1)
	fmt.Printf("%T\n", ss)
	ss = h1
	fmt.Printf("%T\n", ss)
}