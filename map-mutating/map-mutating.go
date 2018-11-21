package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"]) // 42

	m["Answer"] = 48
	fmt.Println(m["Answer"]) // 48

	delete(m, "Answer")
	fmt.Println(m["Answer"]) // 0

	v, ok := m["Answer"]
	fmt.Println("The Value", v, "Present?", ok) // The Value 0 Present? false
}

/*
在映射 m 中插入或修改元素： m[key] = elem
获取元素：elem = m[key]
删除元素：delete(m, key)

通过双赋值检测某个键是否存在：elem, ok = m[key]
若 key 在 m 中，ok 为 true ；否则，ok 为 false。
若 key 不在映射中，那么 elem 是该映射元素类型的零值。
同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

*/
