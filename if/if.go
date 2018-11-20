package main

import (
	"fmt"
	"math"
)

//  if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始不能使用v

	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
