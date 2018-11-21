package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13 // 5*5+12*12取平方根

	// 函数作为入参
	fmt.Println(compute(hypot))    // 5 // 3*3+4*4取平方根
	fmt.Println(compute(math.Pow)) // 81  // 3的4次方
}

/*
函数也是值。它们可以像其它值一样传递。
函数值可以用作函数的参数或返回值。
*/
