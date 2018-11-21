package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	/*
		// fmt.Println(t.S)
		直接访问空指针的成员变量，会有运行错误
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x0 addr=0x0 pc=0x48fe51]
	*/

	if t == nil {
		fmt.Println("<nil>") // 对象为空时也会执行这个接口
		return
	}
	fmt.Println(t.S) // 使用t之前先判空处理了

	// fmt.Println(t)
}

func main() {
	var i I

	var t *T
	i = t       // i指向一个值为空的对象
	describe(i) // (<nil>, *main.T)
	i.M()       // <nil>

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	i.M()       // hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
底层值为 nil 的接口值
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/
