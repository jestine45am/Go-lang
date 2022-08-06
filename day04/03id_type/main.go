package main

import "fmt"

type myint int
type yourint = int
func main(){
	var n myint
	var m yourint
	m = 100
	n = 100
	fmt.Printf("n = %d, n's type is %T\n", n, n)
	fmt.Printf("m = %d, m's type is %T\n", m, m)
	var i byte = 'c'
	var j rune = '中'
	fmt.Printf("%c, %T\n", i, i)
	fmt.Printf("%c, %T\n", j, j)
}