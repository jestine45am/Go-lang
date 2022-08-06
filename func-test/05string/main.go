package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
	/*
	index := strings.IndexRune("helloworld,世界",'世')
	fmt.Println(index)
	f := func(r rune) bool{
		return unicode.Is(unicode.Han,r)
	}
	fmt.Println(strings.IndexFunc("helloworld,世界",f))
	*/
	str := `
			hello world!
			你好世界!号
	`
	f := func(r rune) bool{
		return unicode.Is(unicode.Han,r)
	}
	s := strings.TrimFunc(str,f)
	fmt.Println(s)
}