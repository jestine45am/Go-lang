package main

import "fmt"

func main(){
/*
	var a [4] int
	a = [4]int{1,2,3,4}

	b := [...]int{0,2,3,7,9,33,77}

	c := [5] int{1,5}

	d := [8] int {0:1,6:1,7:9}

	fmt.Printf("Type: %T\t Value: %v\n", a, a)
	fmt.Printf("Type: %T\t Value: %v\n", b, b)
	fmt.Printf("Type: %T\t Value: %v\n", c, c)
	fmt.Printf("Type: %T\t Value: %v\n", d, d)
*/
/*
	citys := [...]string{"beijing","shanghai","guangzhou","shenzheng"}	
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	*/
/*
	citys := [...]string{"beijing","shanghai","guangzhou","shenzheng"}	
	for i,v := range citys {
		fmt.Printf("%v\t%v\n", i, v)
	}
	*/
/*
	a := [3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(a)

*/
/*
	a := [5][4][3]int{
		[4][3]int{
			[3]int{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
			[3]int{10, 11, 12},
		},
		[4][3]int{
			[3]int{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
			[3]int{10, 11, 12},
		},
		[4][3]int{
			{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
			[3]int{10, 11, 12},
		},
		[4][3]int{
			[3]int{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
			[3]int{10, 11, 12},
		},
		[4][3]int{
			[3]int{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
			[3]int{10, 11, 12},
		},
	}
/*
	for i :=0; i < 5; i++{
		for j :=0; j < 4; j++{
			for k := 0; k < 3; k++{
				fmt.Printf("a[%d][%d][%d]=%d\n", i, j, k, a[i][j][k])
			}
		}
	}
	*/
/*
	for _,v1 := range a {
		fmt.Printf("%v\n", v1)
		for _,v2 := range v1 {
			for _,v3 := range v2{
				fmt.Println(v3)
			}
		}
	}
	*/
	/*
	a := [6]int{1,2,3,4,5,6}
	b := a
	b[0] = 99
	fmt.Printf("a:%v\n",a)
	fmt.Printf("b:%v\n",b)
	*/
/*
	a := [5]int {1,3,5,7,8}
	sum := 0
	for i := 0; i < len(a); i++{
		sum += a[i]
	}
	fmt.Println(sum)
	*/
	/*
	a := [5]int {1,3,5,7,8}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a); j++{
			sum := a[i] + a[j]
			if sum == 8 {
				fmt.Printf("%d\t%d\n", i, j)
			}
		}
	}
*/
	var b int
	b = 6
	var a map[int]string
	a = map[int]string{0:"zero", 1:"one"}
	fmt.Printf("%T\t%v\n",a,a)
	fmt.Printf("%T\t%v\n",b,b)
}
