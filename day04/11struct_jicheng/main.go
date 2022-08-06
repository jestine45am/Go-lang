package main

import "fmt"
type animal struct{
	name string
}

type dog struct{
	feet uint8
	animal
}

func (a animal) move(){
	fmt.Printf("%s can move\n", a.name)
}
func (d dog) wang(){
	fmt.Printf("%s has %d feet \n", d.name, d.feet)
}

func main(){
	var d1 = dog{4, animal{"dog"}}
	d1.move()
	d1.wang()
}