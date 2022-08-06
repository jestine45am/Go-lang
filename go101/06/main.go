/*
package main
import "fmt"
func main() {
    // 三个类型不确定常量。它们的默认类型
    // 分别为：int、rune和complex64.
    const X, Y, Z = 2, 'A', 3.12
    var a int = X // 两个类型确定值
	var b int32 = Y
	var z int8 = 127
	o := z + z
	var _ = uint32(32)
	const _ int64 = 1 << uint32(32)
	const y float64 = 
	_ = y
    // 变量d的类型被推断为Y的默认类型：rune（亦即int32）。
    d := X + Z
    // 变量e的类型被推断为a的类型：int。
    e := Y - a
    // 变量f的类型和a及b的类型一样：int。
    f := X * b
    // 变量g的类型被推断为Z的默认类型：complex64。
    g := Z * Y
    // 2 65 (+0.000000e+000+3.000000e+000i)
	fmt.Printf("X: %v, %T\n", X, X)
	fmt.Printf("Y: %v, %T\n", Y, Y)
	fmt.Printf("d: %v, %T\n", d, d)
	fmt.Printf("e: %v, %T\n", e, e)
	fmt.Printf("f: %v, %T\n", f, f)
	fmt.Printf("g: %v, %T\n", g, g)
	fmt.Printf("W: %v, %T\n", W, W)
	fmt.Printf("o: %v, %T\n", o, o)
    // 67 63 130 (+0.000000e+000+1.950000e+002i)
}
*/
package main
//const n = uint(8)
const n uint = uint(8)
var m uint = uint(8)
func main() {
    println(a, b) // 2 0
}
var a = 1 << n / 128
var b = 1 << m / 128