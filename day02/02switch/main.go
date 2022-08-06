package main

import "fmt"

func main(){
	/*
	switch i := 'b'; i {
	case 'g':
			fmt.Println("green")
	case 'r':
			fmt.Println("red")
	case 'b':
			fmt.Println("blue")
	default:
			fmt.Println("None")
	} 
	*/
	var i = 'g'
	switch {
	case i == 'g':
			fmt.Println("green")
			fallthrough
	case i == 'r':
			fmt.Println("red")
	case i == 'b':
			fmt.Println("blue")
	default:
			fmt.Println("None")
	}
}