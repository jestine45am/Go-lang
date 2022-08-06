package main
import "fmt"
const X = 3
const Y = X + X
var a = X
func main() {
    b := Y
	fmt.Printf("X: %T, %v\n", X, X)
	fmt.Printf("Y: %T, %v\n", Y, Y)
	fmt.Printf("a: %T, %v\n", a, a)
	fmt.Printf("b: %T, %v\n", b, b)
}