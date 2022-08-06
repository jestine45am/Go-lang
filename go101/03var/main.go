package main
import "fmt"
var x, y = a+1, 5         // 8 5
var a, b, c = b+1, c+1, y // 7 6 5
var s = -1.23
// var p1 = s
func main(){
	var lang, dynamic = "go", false

	var compiled, announceYear = true, 2009

	var website = "https://golang.org"

	fmt.Printf("lang: %T\n", lang)
	fmt.Printf("dynamic: %T\n", dynamic)
	fmt.Printf("complied: %T\n", compiled)
	fmt.Printf("announceYear: %T\n", announceYear)
	fmt.Printf("website: %T\n", website)
	fmt.Println(a,b,c,x,y)
	p1 := int(s)
	fmt.Println(p1)
	fmt.Println(^int(0)>>63)
	var p2 int = 100
	var p3 int64 = int64(p2)
	fmt.Println(p2, p3)

}