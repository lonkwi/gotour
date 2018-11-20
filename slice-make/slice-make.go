package main

import "fmt"

// func make([]T, len, cap) []T

func main() {
	a := make([]int, 5) // 当容量参数被忽略时，它默认为指定的长度
	printSlice("a", a)  // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
a := make([]int, 5)  // len(a)=5

要指定它的容量，需向 make 传入第三个参数：
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
