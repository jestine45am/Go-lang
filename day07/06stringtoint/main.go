package main

import "fmt"

func StringToInt(str string) int {
	var b = []byte(str)
	var value, sum int
	for i,v := range b{
		value = int(v-'0')
		for j := 0; j < len(b)-i-1; j++{
			value *= 10
		}
		sum += value
	}
	return sum
}

func main(){
	fmt.Println(StringToInt("6789"))
}