package main

import (
	"fmt"
	//"strings"
)
var str string = "hello world!"
var p []byte = []byte(str)

type s1 struct{
	name string
	i int32
}
var s =s1{"Tom",18}

func main(){
	/*
	fmt.Printf("p len:%d\n", len(p))
	fmt.Printf("str len:%d\n", len(str))
	fmt.Printf("str: %q\n", p)
	*/
	/*
	a := "hello world!"
	b := "Hello world!"
	c := "你好go语言"
	fmt.Println(strings.Compare(a,b))
	fmt.Println(strings.Compare(a,a))
	fmt.Println(strings.Compare(b,a))
	fmt.Println(strings.Contains(a,"o "))
	fmt.Println(strings.ContainsRune(c,'点'))
	fmt.Println(strings.ContainsAny(a,"x y"))
	fmt.Println(strings.Count(a,"l"))
	fmt.Println(strings.Count("fivevev", "vev"))

	fmt.Println(len(c))
	fmt.Println(strings.Count(a,""))
	*/

//	fmt.Printf("%s\f",str)
	const MaxInt = (^uint(0) >> 63)
	fmt.Println(MaxInt)
	fmt.Printf("%d\n", MaxInt)
}




















