package main

import "fmt"
// 猜测 空接口的类型
func assign(i interface{}){
	fmt.Printf("value: %v\ttype:%T\n",i,i)
	switch v := i.(type){
	case string: fmt.Printf("%v transfer to type string's value:%v\n",i,v)
	case int: fmt.Printf("%v transfer to type int's value:%v\n",i,v)
	case int64: fmt.Printf("%v transfer to type int64's value:%v\n",i,v)
	case int32: fmt.Printf("%v transfer to type int32's value:%v\n",i,v)
	case bool: fmt.Printf("%v transfer to type bool's value:%v\n",i,v)
	case float32: fmt.Printf("%v transfer to type float32's value:%v\n",i,v)
	case float64: fmt.Printf("%v transfer to type float64's value:%v\n",i,v)

	}
}
func main(){
	assign("hello world")
	assign(100)
	assign(true)
	assign(3.26)
	assign(int32(33))
	var i = 3.21
	var a = float32(i)
	assign(a)
}
/*
func assign(i interface{}){
	fmt.Printf("%T\t%v\n",i,i)
	str,ok := i.(string)
	if !ok {
		fmt.Println("err")
	}else {
		fmt.Println(str)
	}
}
func main(){
	var i = "hello world!"
	assign(i)
}
*/
