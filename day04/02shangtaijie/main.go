package main
import "fmt"
// 走台阶
// 可以一次走一个台阶
// 也可以一次走两个台阶
func f(n uint)(m uint){
	if n == 1 {
		m = 1
	}else if n == 2 {
		m = 2
	}else {
		m = f(n - 1) + f(n - 2)
	}
	return 
}
func main(){
	x := f(10)
	fmt.Println(x)
}