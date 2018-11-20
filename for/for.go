package main

import "fmt"

// Go 只有一种循环结构：for 循环。

// 初始化语句：在第一次迭代前执行
// 条件表达式：在每次迭代前求值
// 后置语句：在每次迭代的结尾执行

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 0
	for i := 10; i < 10; i++ { // 条件表达式第一次就为false
		sum2 += i
	}
	fmt.Println(sum2) // 值为0, 每次迭代前检查条件，可以一次都不执行

	// 初始化语句和后置语句是可选的。
	sum3 := 1
	for sum3 < 1000 { // 只有条件表达式
		sum3 += sum3
	}
	fmt.Println(sum3)

	// 可以去掉分号，因为 C 的 while 在 Go 中叫做 for。。
	sum4 := 1
	for sum4 < 1000 { // 只有条件表达式，并且省略了前后分号
		sum4 += sum4
	}
	fmt.Println(sum4)

	// 无限循环
	j := 1
	for { // 不带条件
		j++
		if j > 1000 {
			fmt.Println(j)
			break
		}
	}
}
