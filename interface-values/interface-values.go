package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T) // 具体底层类型(*main.T)的具体值(&{Hello})
	i.M()       // Hello

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M()       // 3.141592653589793
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
接口值
在内部，接口值可以看做包含值和具体类型的元组：
(value, type)
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。
*/
