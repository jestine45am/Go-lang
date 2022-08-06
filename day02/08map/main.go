package main

import (
	"fmt"
)

func main()  {
	/*
	var m1 = make(map[string]int, 10)
	fmt.Println(m1 == nil)
	m1["hello world"] = 0
	m1["ni hao"] = 1
	fmt.Println(m1)

	var m2 map[string]int
	fmt.Println(m2 == nil)
*/	
/*
	m1 := make(map[string]int, 10)
	m1["one"] = 1
	m1["two"] = 2
	fmt.Println(m1)
	fmt.Println(m1["one"])
	fmt.Println(m1["two"])
	fmt.Println(m1["three"]) // 查看未定义的值
*/
/*
	// 判断某一个键是否存在
	
	m1 := make(map[string]int, 10)
	m1["one"] = 1
	m1["two"] = 2
	value, ok := m1["two"]
	if !ok {
		fmt.Println("This key is not exist")
	} else {
		fmt.Println(value)
	}
	*/
	/*
	// map 的遍历
	m1 := make(map[string]int, 10)
	m1["one"] = 1
	m1["two"] = 2
	/*

	for key, value := range m1 {
		fmt.Println(key, value)
	}
	for _, value := range m1 {
		fmt.Println(value)
	}
	*/
	/*
	// 删除指定的键值对
	m1 := make(map[string]int, 10)
	m1["one"] = 1
	m1["two"] = 2
	delete(m1, "two")
	fmt.Println(m1)
	
	*/
	// 声明元素类型为 map 的切片
	//方法一：
/*
	var s1 = make([]map[int]string, 0, 100)
	fmt.Println(s1)
	var m1 = make(map[int]string, 10)
	m1[1] = "one"
	m1[2] = "two"
	m1[3] = "three"
	
	var m2 = make(map[int]string, 2)
	m2[1] = "one"
	m2[2] = "two"
	m2[3] = "three" 
	s1 = []map[int]string{m1,m2} 

	fmt.Println(s1)
	*/
	/*

	// 方法二:
	var s1 = make([]map[int]string, 10, 10)  // 初始化，并创建了一个底层数组
	for i := 0; i < 10 ; i++ {
		s1[i] = make(map[int]string,1)  // 底层数组的元素类型是 map，map 作为元素，需要被赋值，那么 map 本身也需要初始化
	}									// 否则就是向 nil 赋值，会发生错误
	s1[0][0] = "zero"
	s1[1][1] = "one"
	s1[2][2] = "two"
	fmt.Println(s1)
	*/
	
	// 值为切片类型的 map
/*
	var m1 = make(map[int][]string, 10)
	var s1 = []string{"one","two","three"}
	var s2 = []string{"one","two","three"}
	fmt.Println(s1)
	m1[0] = s1
	m1[1] = s2
*/
/*
	var m1 = make(map[int][]string,10)
	m1[0] = []string{"one", "two", "three"}
	m1[1] = []string{"four", "five"}
	fmt.Println(m1)
	*/
/*
	// 练习：统计 "how do you do" 中各个单词出现的次数
	// 方法一:
	var s = []string{"how", "do", "you", "do"}
	fmt.Println(s)
	var m = make(map[string]int, 100)
	for i := 0; i < len(s); i++{
		m[s[i]] = 1
	}
	for i :=0; i < len(s); i++ {
		for j := 0; j < len(s) && j != i; j++{
			if s[i] == s[j] {
				m[s[i]] += 1
			}
		}
	}

	fmt.Println(m)
	// 方法二:


	s1 := "how do you do"
	s2 := strings.Split(s1, " ")
	m1 := make(map[string]int, 10)
	for _, w :=range s2{
		_, ok := m1[w]
		if !ok {
			m1[w] = 1
		}else {
			m1[w] +=1
		}

	}
	fmt.Println(m1)
	*/
/*
	m1 := make(map[string]int,1)
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	fmt.Println(m1)
	*/
	// 回文判断
	// 上海自来水来自海上
	a := "a上海自来水来自海上a"
	s := make([]rune, 0, len(a))
	for _, c := range a {
		s = append(s, c)
	}
	for i := 0; i < len(s)/2; i++{
		if s[i] != s[len(s)-1-i]{
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}



