package main
import "fmt"

type human struct{
	name string
	age int
	gender string
	hobby []string
}
func main(){
	/*
	var tom human
	tom.name = "tom"
	tom.age = 18
	tom.gender = "男"
	tom.hobby = []string{"swimming","pingpong","football"}
	fmt.Println(tom)
*/
/*
	var jim struct{
		name string
		age int
		gender string
	}
	jim.name = "jim"
	jim.age = 19
	jim.gender = "男"
	*/
	/*
	// 匿名结构体
	jim = struct{name string; age int; gender string}{"tom",19, "nan"}
	fmt.Printf("type: %T\nvalue: %v", jim, jim)
*/
/*
	// 结构体指针
	var h1 = new(human)
	(*h1).name = "John"
	h1.name = "John"
	fmt.Println(h1.name)
	fmt.Println((*h1).name)
	*/
	 // 结构体赋值
	 type person struct{
		 name string
		 age int
	 }
	 var p2 = person{"Tom", 18}
	 fmt.Println(p2)
	 var p1 = human{
		 name: "John",
		 age: 18,
		 gender: "male",
		 hobby: []string{"football","running"},
	 }
	 fmt.Printf("%p\n", &p1.name)
	 fmt.Printf("%p\n", &p1.age)
	 fmt.Printf("%p\n", &p1.gender)
}