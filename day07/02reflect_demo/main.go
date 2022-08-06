package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
}
func reflectValue(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Printf("value:%v\n", v)
}
type human struct{
	name string
	age int
}
func main(){
	var a float32 = 3.14
	var b int64 = 100
	var h human = human{"Tom", 18}
	reflectType(a)
	reflectValue(a)
	reflectType(b)
	reflectValue(b)
	reflectType(h)
	reflectValue(h)
}