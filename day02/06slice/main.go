package main

import "fmt"

func main(){
	/*
	var s1 []int
	var s2 []string
	s1 = []int{1,3,5,7,9}
	s2 = []string{"hello","world","!"}
	fmt.Println(s1, len(s1))
	fmt.Println(s2, len(s2))
	*/
/*
	var s1 []int
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = []int{1,3,5,7}
	s2 = []string{"hello","world"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println("s1: len(s1)", len(s1))
	fmt.Println("s2: len(s2)", len(s2))
	fmt.Println("s1: cap(s1)", cap(s1))
	fmt.Println("s2: cap(s2)", cap(s2))
*/
/*
	a := [...]int{1,3,5,7,9,11}
	s := a[1:4] // [3,5,7]
	s1 := s[1:3] // [5,7]

	fmt.Println(s)
	fmt.Println(s1)
	
	fmt.Println("len(s):",len(s),"cap(s):",cap(s))
	fmt.Println("len(s1):",len(s1),"cap(s1):",cap(s1))
	*/
	// make函数创建切片
	/*
	s1 := make([]int, 5, 10)
	fmt.Println(s1, len(s1), cap(s1))
	*/
/*
	// 切片的赋值
	s1 := [] int {1, 2, 3, 4, 5}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)
	*/

/*
	s1 := [] int {1, 2, 3, 4, 5}
	for i := 0; i < len(s1); i++{
		fmt.Printf("s1[%d]=%d\n", i, s1[i])
	}
*/
/*
	s1 := [] int {1, 2, 3, 4, 5}
	for i,v := range s1 {
		fmt.Printf("s1[%d]=%v\n", i, v)
	}
	*/

/*
	// 切片的扩容


	s1 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("s1: %v\t len(s1): %d\tcap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 99)
	fmt.Printf("s1: %v\t len(s1): %d\tcap(s1):%d\n", s1, len(s1), cap(s1))
	ss := []int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	s1 = append(s1, ss...)
	fmt.Printf("s1: %v\t len(s1): %d\tcap(s1):%d\n", s1, len(s1), cap(s1))
*/
	// 切片的拷贝
/*
	s1 := [] int {1, 2, 3, 4, 5}
	s2 := s1   // 切片的赋值
	s3 := make([]int, 2, 20)
	copy(s3, s1)  // 切片的拷贝
	fmt.Printf("s1=%v\ts2=%v\ts3=%v\n", s1, s2, s3)
	s1[0] = 100
	fmt.Printf("s1=%v\ts2=%v\ts3=%v\n", s1, s2, s3)

*/
/*
	// 删除切片中指定索引的元素

	s1 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n",s1, len(s1), cap(s1))
	s1 = append(s1[:1],s1[2:]... )
	fmt.Println(s1)
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n",s1, len(s1), cap(s1))
	*/
	/*
	// 切片练习
	var s = make([]string, 5, 10)
	fmt.Println(s)
	for i := 0; i < 10; i++{
		s = append(s, fmt.Sprintf("%v", i))
	}
	fmt.Println(s, len(s), cap(s))
	*/
	/*
	var s []int
	s = append(s, 10)
	fmt.Println(s)
	*/
	m1 := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println(m1)
}