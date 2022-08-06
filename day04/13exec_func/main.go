package main

import "fmt"
/*
创建学生管理系统
功能：
	1. 查看所有学生信息
	2. 新增学生信息
	3. 删除学生信息
	4. 退出
*/	
// student 结构
type student struct{
	id uint
	name string
	age uint
}
// map 类型
var m1 = make(map[uint]student, 100)
// 打印所有的 student
func showAllStudent(){
	for _,k := range m1{
		fmt.Printf("%v\n",k)
	}
}
// 添加 student 结构变量
func addStudent(){
	var id uint
	var name string
	var age uint
	fmt.Print("Please enter the student's name:")
	fmt.Scanf("%s\n",&name)
	fmt.Print("Please enter the student's id:")
	fmt.Scanf("%d\n",&id)
	fmt.Print("Please enter the student's age:")
	fmt.Scanf("%d\n",&age)
	var s1 = student{id,name,age}
	m1[id] = s1
}

// 删除 m1 中的student
func deleteStudent(){
	var id uint
	fmt.Print("Please enter the student's id you want to delete:")
	fmt.Scanf("%d\n",&id)
	delete(m1, id)
}

func main(){
	// 打印提示信息，等待用户选择
	for {
		fmt.Println("欢迎来到学生管理系统！")
		fmt.Println( `
			1. 查看所有学生信息
			2. 新增学生信息
			3. 删除学生信息
			4. 退出
			`)
		var choice uint8
		fmt.Print("请选择需要进行的操作：")
		fmt.Scanf("%d\n",&choice)
		switch choice{
		case 1: showAllStudent()
		case 2: addStudent()
		case 3: deleteStudent()
		case 4: goto out
		default: fmt.Println("please enter the correct num!")
		}
	}
	out:
		fmt.Println("GoodBye!")	
}