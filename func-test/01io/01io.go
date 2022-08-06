package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	
	f1,err := os.Open("./test.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	tmp := [128]byte{}
	for{
		n1,err01 := f1.Read(tmp[:])
		if err01 != nil{
			fmt.Println(err01)
			return
		}
		if err01 == io.EOF{
			return
		}
		fmt.Println(string(tmp[:n1]))

		f2,err02 := os.OpenFile("./wr_test.txt", os.O_CREATE|os.O_APPEND, 0644)
		if err02 != nil{
			fmt.Println(err02)
			return
		}
		_,err03 := f2.Write(tmp[:n1])
		if err03 != nil{
			fmt.Println(err03)
			return 
		}

	}

}