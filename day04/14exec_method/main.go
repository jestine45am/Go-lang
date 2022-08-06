package main

import (
	"fmt"
	"os"
)

// student 结构保存 student 的属性信息
type student struct{
	id uint
	name string
	age uint
}
// 创建一个 map 来存储数据
var m = make(map[uint]student, 100)

// studentMgr 结构体用来存储所有的学生数据
type studentMgr struct{
	allstudent map[uint]student
}

// 查看所有学习信息的方法
func (s *studentMgr) showAllStudent(){
	fmt.Println("学生信息如下：")
	for _,v := range s.allstudent{
		fmt.Printf("学生学号：%d 学生姓名：%s 学生年龄：%d\n", v.id, v.name, v.age)
	}

}
// 添加学生信息的方法
func (s *studentMgr) addStudent(){
	var id, age uint
	var name string
	fmt.Println("正在进行添加学生信息的操作！")
	fmt.Print("请输入学生的学号:")
	fmt.Scanf("%d\n",&id)
	// 检查输入的学号是否已经存在，如果存在则打印学生信息，并退出
	// 如果不存在，则进行添加操作
	v,ok := s.allstudent[id]
	if ok {
		fmt.Println("该学生信息已经存在")
		fmt.Printf("学生学号：%d 学生姓名：%s 学生年龄：%d\n", v.id, v.name, v.age)
		return
	}else {
		fmt.Print("请输入学生的姓名:")
		fmt.Scanf("%s\n",&name)
		fmt.Print("请输入学生的年龄:")
		fmt.Scanf("%d\n",&age)
		s.allstudent[id] = student{id,name,age}
		fmt.Println("学生信息添加成功！")
	}
}
// 修改学生信息的方法
func (s *studentMgr) editStudent(){
	// 提示用户输入学号
	// 查找学生学号是否存在，如果存在，则进行修改操作，如果不存在，在提示用户不存在，退出
	var (
		id uint
		name string
		age uint
	)

	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	v,ok := s.allstudent[id]
	if !ok {
		fmt.Println("该学生不存在")
		return
	}else {
		fmt.Println("该学生信息存在，信息如下：")
		fmt.Printf("学生学号：%d 学生姓名：%s 学生年龄：%d\n", v.id, v.name, v.age)
		fmt.Print("请输入学生的姓名:")
		fmt.Scanf("%s\n",&name)
		fmt.Print("请输入学生的年龄:")
		fmt.Scanf("%d\n",&age)
		s.allstudent[id] = student{id,name,age}
		fmt.Println("学生信息修改成功")
	}
}
// 删除学生信息的方法
func (s *studentMgr) deleteStudent(){
	// 提示用户输入需要删除的学生学号
	// 检查输入的学号是否存在，如果存在，则执行删除操作，如果不存在，则不执行任何操作
	var id uint
	fmt.Print("请输入需要删除的学生学号:")
	fmt.Scanf("%d\n",&id)
	_,ok := s.allstudent[id]
	if !ok {
		fmt.Println("该学生信息不存在")
		return
	}else {
		delete(s.allstudent,id)
		fmt.Println("学生信息删除成功！")
	}
}

func main(){
	var sm = studentMgr{m}
	for {
	fmt.Println("欢迎来到学生管理系统！")
	fmt.Println(
		`
		1. 查找所有的学生信息
		2. 添加学生信息
		3. 删除学生信息
		4. 修改学生信息
		5. 退出
		`)
	var choice uint8
	fmt.Print("请选择需要进行的操作：")
	fmt.Scanf("%d\n",&choice)
	switch choice {
	case 1: sm.showAllStudent() 
	case 2: sm.addStudent() 
	case 3: sm.deleteStudent() 
	case 4: sm.editStudent()
	case 5: os.Exit(1)
	default: fmt.Println("请输入正确的序号！")
	}
}
}