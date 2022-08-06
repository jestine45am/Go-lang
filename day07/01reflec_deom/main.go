package main

import (
	"encoding/json"
	"fmt"
)

type human struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	str := `{"name":"zbc","age":18}`
	b := []rune(str)
	fmt.Printf("%T\n",b)
	var h human
	json.Unmarshal([]byte(str), &h)
	fmt.Println(h.Name, h.Age)
	v,err := json.Marshal(h)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(v))
}