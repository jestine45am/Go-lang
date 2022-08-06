package main

import (
	"fmt"
	calc "github.com/fockpmov34/day05/07package/calc"
)

func init(){
	fmt.Println(i,pi)
}

var i = 100
const pi = 3.14


func main(){
	sum := calc.Add(1,7)
	fmt.Println("sum = ",sum)
}