package main

import "fmt"

func main(){
/*
	a := 5
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a ++
	b --
	fmt.Println(a)
	fmt.Println(b)
	*/
	/*
	a := 5
	b := 2
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	*/
/*
	age := 22
	fmt.Println(age > 18 && age < 19)
	fmt.Println(age < 18 || age < 60)
	fmt.Println(!(age < 18))
	*/
/*
	a := 0b00000001
	b := 0b00010000
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 4)
	fmt.Println(b >> 4)
*/
/*
	var i = int8(1)
	fmt.Println(i << 10)
	*/
	var i int
	i = 5
	fmt.Println(i) // 5
	i += 2
	fmt.Println(i) // 7
	i -= 2
	fmt.Println(i) // 5
	i *= 2
	fmt.Println(i) // 10
	i /= 2
	fmt.Println(i) // 5
	i %= 2
	fmt.Println(i) // 1
	i &= 2
	fmt.Println(i) // 0
	i |= 2
	fmt.Println(i) // 2
	i ^= 2
	fmt.Println(i) // 0
	i <<= 2
	fmt.Println(i) // 0
	i >>= 2
	fmt.Println(i) // 0
}