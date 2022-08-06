package main
import "fmt"

// 函数是一直数据类型，包含函数的参数和返回值
// 函数可以作为另一个函数的参数，进行参数传递
// 函数也可以作为另一个函数的返回值

// 一般函数
func f1(){
	fmt.Println("hello world!")
}

func f2()int{
	return 10
}
// 函数作为参数
func f3(x func()int){
	ret := x()
	fmt.Println(ret)
}

func ff(x, y int)int{
	return x + y
}
// 函数作为返回值
func f4(func()int) func(int,int)int{
	return ff
}
func main(){
	a := f1
	b := f2
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	f3(f2)
	f7 := f4(f2)
	fmt.Printf("%T\n",f7)
	// 函数名本身是一个地址，表示函数的入口
	fmt.Println(ff)
}