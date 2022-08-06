package main

import "fmt"
type everyType interface{}
func show(i everyType){
	fmt.Println(i)
}
func main(){
	var a = 100
	show(a)
	var b = []string{"hello","world"}
	show(b)
}

/*
func main(){
	//interface{} 用于map的value类型
	m1 := make(map[string]interface{}, 20)
	m1["name"] = "Tom"
	m1["age"] = 18
	m1["married"] = false
	m1["hobby"] = []string{"football","swimming","running"}
	fmt.Println(m1)

}
*/