package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big)) 编译错误 constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Printf("type of Small: %T\n", Small)
	// fmt.Printf("type of Big: %T\n",Big) // 编译错误 constant 1267650600228229401496703205376 overflows int
}
