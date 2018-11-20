package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(335)
	fmt.Println("my favorite number is", rand.Intn(10))
}
