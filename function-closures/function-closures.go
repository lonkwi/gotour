package main

import "fmt"

/*
Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
该函数可以访问并赋予其引用的变量的值，换句话说，该函数被“绑定”在了这些变量上。
例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"Loop ", i, " ->",
			pos(i), // 闭包pos的sum值每个循环都往上增长
			neg(-2*i),
		)
	}
}

/*
Loop  0  -> 0 0
Loop  1  -> 1 -2
Loop  2  -> 3 -6
Loop  3  -> 6 -12
Loop  4  -> 10 -20
Loop  5  -> 15 -30
Loop  6  -> 21 -42
Loop  7  -> 28 -56
Loop  8  -> 36 -72
Loop  9  -> 45 -90
*/
