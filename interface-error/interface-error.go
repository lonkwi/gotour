package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// *MyError实现接口Error
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// 函数返回一个error接口
func run() error {
	// 返回的error接口由&MyError实现
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err) // err接口返回字符串，可以直接打印
	}
}

/*
Go 程序使用 error 值来表示错误状态。
与 fmt.Stringer 类似，error 类型是一个内建接口：
type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）

通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示失败。
*/
