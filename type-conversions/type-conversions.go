package main

import (
	"fmt"
	"math"
)

// 表达式 T(v) 将值 v 转换为类型 T。
// 与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换
// 没有显示转换的会报编译错误
// 示例： cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}
