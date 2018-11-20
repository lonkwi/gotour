package main

import "fmt"

// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128

func main() {
	v := 0.867 + 0.5i // 3.142 // 42
	j := v
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("j is of type %T\n", j)
}
