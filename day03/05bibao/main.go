package main
import (
	"fmt"
	"strings"
)


// 闭包
/*
func f1(func()){
	fmt.Println("This is f1")
}

func f2(x,y int)int{
	return x+y
}
func f3(f func (int,int)int, x,y int) func(){
	return func(){
		f(x,y)
	}
}
func main(){
	ret := f3(f2,100,200)
	fmt.Printf("%T\n",ret)
	f1(ret)
}
*/
/*
func f1(x int) func(int)int{
	return func(y int)int {
		x += y
		return x
	}
}
func main(){
	ret := f1(100)
	ret2 := ret(200)
	fmt.Println(ret2)
}
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(jpgFunc("test.jpg")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
