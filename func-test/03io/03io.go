package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Golang")
	p := make([]byte, 10)
	n,err := reader.ReadAt(p,2)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(p[1:n]))
	fmt.Println("读取的字节数是：", n)

	
}