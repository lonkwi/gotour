package main

import (
	"fmt"
	"math"
)

// 定义接口Abser
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser // 申明一个接口变量 // 类比C++的基类？
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex 实现了 Abser
	fmt.Println(a.Abs())

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v // 编译不过
}

type MyFloat float64

// MyFloat实现了Abs函数；实现了接口Abser
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// *Vertex实现了Abs函数；实现了接口Abser
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
