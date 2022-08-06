package main

import "fmt"

const (
    X float32 = 3.14
    Y                // 这里必须只有一个标识符
    Z                // 这里必须只有一个标识符
    A, B = "Go", "language"
    C, _
    // 上一行中的空标识符是必需的（如果
    // 上一行是一个不完整的常量描述）。
)

func main(){
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}