package main

import "fmt"

/*
const (
	pi = 3.1415926
	userNotExistError = 100000

) 
*/

func main()  {
	/*
	s1 := "Hello沙河小王子"
	//fmt.Printf("%s\n", s1)
	s2 := []rune(s1)
	for i := 0; i < len(s1); i++{

		fmt.Printf("%c\n",s2[i])
		
	}

	var age = 19
	if age > 18 {
		fmt.Println("已经成年了")
	} else if age > 65 {
		fmt.Println("已经退休了")
	} else {
		fmt.Println("好好上班挣钱吧！")
	}

	var  i = 0
	for ; i < 10; {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("无限循环")
	}

	s := "hello"
	fmt.Println(len(s))
	for _,v := range s{
		fmt.Printf("%c\n",v)

	}

	var f1 = 3.14
	fmt.Printf("%T\n", f1)
	*/
/*
	var s1 = "hello world"
	//var s2[] byte = []byte(s1)
	var s2 = []byte(s1)
	s2[0] = 'H'
	for i:=0; i<len(s1); i++ {
		fmt.Printf("%c\n",s2[i])
	}
	*/
/*
	var c1 byte = 'c'
	var c2 rune = 'd'
	fmt.Printf("%c, %T\n", c1, c1)
	fmt.Printf("%c, %T\n", c2, c2)
	*/
/*
	var c1 bool = true
	fmt.Printf("%T %v\n", c1, c1)
	*/

	/*
	c5 := "你"
	c6 := "H"
	fmt.Printf("%s %d\n", c5, len(c5))
	fmt.Printf("%s %d\n", c6, len(c6))
*/
	num2 := 0b11111111
	num8 := 077
	num10 := 55
	num16 := 0xff
	fmt.Printf("%d\n",num2)
	fmt.Printf("%d\n",num8)
	fmt.Printf("%d\n",num10)
	fmt.Printf("%d\n",num16)
/*
	const (
		num1 = 100
		num2
		num3
	)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
*/
}
