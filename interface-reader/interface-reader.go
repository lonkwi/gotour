package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// NewReader函数返回一个Reader结构体，输入参数是string
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// Reader结构体实现了io/Read接口
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
