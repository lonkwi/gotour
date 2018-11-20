package main

import "fmt"

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。 例如下面的sum
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
