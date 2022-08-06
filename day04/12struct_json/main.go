package main

import (
	"encoding/json"
	"fmt"
)

// 把 Go 语言中的结构体变量转换成 JSON 格式
// 把 JSON 格式的字符串，转换成 Go语言中能够识别的结构体变量

type human struct{
	Name string `json:"name" ini:"name" db:"name"`
	Age int	`json:"age" ini:"age" db:"age"`
}
func main(){
	// 正向解析
	h1 := human{"Tom",18} 
	v,err := json.Marshal(h1)
	if err != nil{
		fmt.Printf("Marshal is failed, err: %v\n",err)
		return
	}
	fmt.Printf("%v\n", string(v))

	// 反向解析
	str := `{"name":"Jim","age":18}`
	var h2 human
	json.Unmarshal([]byte(str), &h2)
	fmt.Printf("%#v\n", h2)
}