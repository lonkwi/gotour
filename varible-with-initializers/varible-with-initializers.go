package main

import "fmt"

// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
