package main
import "fmt"

func f1(){
	defer func(){
		err := recover()
		fmt.Println(err)
	}()
	panic("error!!!!!")
	fmt.Println("f1")
}

func f2(){
	fmt.Println("f2")
}
func main(){
	f1()
	f2()
}