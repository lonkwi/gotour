package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I // 接口 i 为空
	describe(i)
	i.M() // 运行错误，不晓得调用哪个具体类的函数
} /*
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x0 pc=0x48fad2]
*/

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
nil 接口值
nil 接口值既不保存值也不保存具体类型。
为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
*/
