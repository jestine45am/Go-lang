package main

import (
	//	"fmt"
	"fmt"
	"reflect"
)

type student struct{
	Name string `json:"name"`
	Age int 	`json:"age"`
}
// reflectValueSet ...
func reflectValueSet(x interface{}){
	v := reflect.ValueOf(x)
	stu := student{"Jim", 10}
	if v.Elem().Kind() == reflect.String {
		v.Elem().SetString(stu.Name)
	} else if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(int64(stu.Age))
	}
}
	
func main(){
	var s = student{"Tom", 18}
/*	
	// 通过反射获取变量的信息
//	fmt.Printf("TypeName: %v\tTypeKind: %v\n", reflect.TypeOf(s).Name(), reflect.TypeOf(s).Kind()) // student  struct
//	fmt.Printf("Value: %v\n", reflect.ValueOf(s))
	fmt.Println(reflect.TypeOf(s).NumField())
	fmt.Println(reflect.TypeOf(s).Field(1))

	// 通过反射修改变量的信息
	reflectValueSet(&s.Name)
	reflectValueSet(&s.Age)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s).FieldByName("Age"))

	fmt.Println(reflect.ValueOf(s).NumField())
	fmt.Println(reflect.ValueOf(s).Field(1))
	fmt.Println(reflect.ValueOf(s).FieldByName("Name"))
	v := reflect.ValueOf(s.Name)
	fmt.Printf("%v\n", v)
	*/
	for i := 0; i < reflect.TypeOf(s).NumField(); i++ {
		FieldStr := reflect.TypeOf(s).Field(i)
		fmt.Println(FieldStr.Name, FieldStr.Tag.Get("json"))
	}
}
