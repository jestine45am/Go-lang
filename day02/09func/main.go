package main

import "fmt"

func sum (i int,j int)(s int){
	s = i + j
	return s
}
func f1(i int, j int)(k int){
	k = i + j
	return k
}
func f2(){
	fmt.Println("没有参数，也没有返回值")
}
func f3(i int, j int){
	fmt.Printf("%d + %d = %d",i, j, i+j)
}
func f4()(int){
	return 666
}
func f5()(int, string){
	return 1, "one"
}
func f6(x,y int)(o,p int){
	o = x + y
	p = x - y
	return
}
func f7(x string, y... int){
	fmt.Println(x)
	fmt.Println(y)
}
func main(){
	f7("one")
}