package main

import (
	"fmt"
	"strconv"
)

func main(){
	var str = "10009"
	retint,_ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v\t%T\n", retint, retint)

	retint1,_ := strconv.Atoi(str)
	fmt.Printf("%v\t%T\n", retint1, retint1)

	var boolstr = "false"
	retbool,_ := strconv.ParseBool(boolstr)
	fmt.Printf("%v\t%T\n", retbool, retbool)

	var floatstr = "3.1415926"
	retfloat,_ := strconv.ParseFloat(floatstr, 64)
	fmt.Printf("%v\t%T\n", retfloat, retfloat)

}