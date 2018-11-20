package main

import "fmt"

type Vertex struct {
	X, Y int // 可以同一行写
}

var (
	v1 = Vertex{1, 2}  // has ype Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

/*
结构体文法通过直接列出字段的值来新分配一个结构体。
使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
特殊的前缀 & 返回一个指向结构体的指针。
*/
