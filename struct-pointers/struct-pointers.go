package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// p是指向v的指针
	(*p).X = 1e4 //1e4: 10000, 1e5: 100000 1e9: 1000000000
	// 指针改成员值可以直接用.号引用
	p.Y = 666

	fmt.Println(v)
}

/*
如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
不过这么写太啰嗦了，所以语言也
允许我们使用隐式间接引用，直接写 p.X 就可以。
*/
