package main

import "fmt"
/*
func deferdemo(){
	fmt.Println("start")
	defer fmt.Println("first define, last output")
	defer fmt.Println("last define, first output")
	fmt.Println("end")
}
func f1()int{
	x := 5
	defer func(){
		x++
	}()
	return x
}

func f2()(x int){
	defer func(){
		x ++
	}()
	return 5
}

func f3()(y int){
	x := 5
	defer func(){
		x++
	}()
	return x
}
func f4()(x int){
	defer func(x int){
		x++
	}(x)
	return 5
}
func main()  {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}