package main

import "fmt"

func main() {
	// 指定了 !!! 零个方法 !!! 的接口值被称为 *空接口：*
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
