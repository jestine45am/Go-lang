package main

import (
	"fmt"
	"io"
	"os"
)

func readfrom(reader io.Reader, num int)([]byte, error){
	p := make([]byte, num)
	n,err := reader.Read(p)	
	if n > 0 {
		return p, nil
	}
	return p, err
}

func main(){
	data, err := readfrom(os.Stdin, 10) 
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}