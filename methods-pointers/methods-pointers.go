package main

import (
	"fmt"
	"math"
)

// 类比C语言的函数入参，传值，还是传指针

type Vertex struct {
	X, Y float64
}

// 若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）
// 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)       // 30 40
	fmt.Println(v.Abs()) // 50
}
