package main

import "fmt"

func main()  {
		
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			if j == 3{
				break
			}
			fmt.Printf("%d\t%d\n", i, j)
		}
	}
	/*
	newThread:
		fmt.Println("over")
		*/
	
}