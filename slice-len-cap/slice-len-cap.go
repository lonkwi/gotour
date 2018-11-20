package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6, cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	// 只限制切片的结束位置，容量cap没变化
	s = s[:0]
	printSlice(s) // len=0, cap=6 []

	// Extend its length
	// 只限制切片的结束位置，容量cap没变化
	s = s[:4]
	printSlice(s) // len=4, cap=6 [2 3 5 7]

	// 切片结束位置不能超过容量
	// s = s[:7] // 运行错误panic: runtime error: slice bounds out of range
	// printSlice(s)

	// Drop its first two values
	// 改变起始元素后，容量cap变化
	s = s[2:]
	printSlice(s) // len=2, cap=4 [5 7]

}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}

/*
切片拥有 长度 和 容量。
切片的长度就是它所包含的元素个数。
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
*/
