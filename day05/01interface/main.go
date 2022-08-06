package main

import (
	"fmt"
)
// 声明一个结构体 cat
type cat struct{

}
// 声明一个结构体 dog
type dog struct{

}
// 声明一个结构体 human
type human struct{}

// 声明一个接口 speaker, 其方法是 speak()
type speaker interface{
	speak()
}

// 定义一个针对 cat 结构体的方法 speak()
func (c cat) speak(){
	fmt.Println("喵喵喵~")
}

// 定义一个针对 dog 结构体的方法 speak()
func (d dog) speak(){
	fmt.Println("汪汪汪~")
}

// 定义一个针对 human 结构体的方法 speak()
func (h human) speak(){
	fmt.Println("啊啊啊~")
}

// 定义一个针对 speaker 类型的方法函数
func f(x speaker){
	x.speak()
}

func main(){
	var (
		c1 cat
		d1 dog
		h1 human
	)
	f(c1)
	f(d1)
	f(h1)
}