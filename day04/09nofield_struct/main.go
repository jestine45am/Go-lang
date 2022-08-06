package main
import "fmt"
/*
type student struct{
	name string
	age int
}
*/
type student struct{
	string
	int
}
func main(){
	var s1 = student{"Tom", 18}
	fmt.Println(s1.string)
	fmt.Println(s1.int)
}