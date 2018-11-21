package main

import (
	"fmt"
	"math"
)

/*
！！
接收者为指针时，可以接收指针和值
接受者为值时，也可以接受指针和值。(接收者的成员值始终不会被修改）

函数不行，类型必须匹配
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
	// fmt.Println(AbsFunc(p)) // 编译错误cannot use p (type *Vertex) as type Vertex in argument to AbsFunc
}

/*
接受一个值作为参数的函数必须接受一个指定类型的值：
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！

而以值为接收者的方法被调用时，接收者既能为值又能为指针：
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
*/
