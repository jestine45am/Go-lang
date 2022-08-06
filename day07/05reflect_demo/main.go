package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
/*
func LoadStu(name string, age int, v1, v2 interface{}){
	reflect.ValueOf(v1).Elem().SetString(name)
	reflect.ValueOf(v2).Elem().SetInt(int64(age))
}

func main(){
	var s = student{}
	fmt.Println(&s, &s.Name, &s.Age)
	LoadStu("tom", 18, &s.Name, &s.Age)
	fmt.Println(s)
}
*/
func LoadStu(v interface{}){
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	for i := 0; i < t.Elem().NumField(); i++{
		f := t.Elem().Field(i).Name
		sValue := value.Elem().FieldByName(f)
		sPtr := sValue.Addr()
		switch f {
		case "Name": sPtr.Elem().SetString("Jim")
		case "Age": sPtr.Elem().SetInt(int64(99))
		}

	}
}

func main(){
	var s = student{}
	LoadStu(&s)
	fmt.Println(s)

}