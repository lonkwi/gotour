package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	// 切片可包含任何类型，甚至包括其它的切片。
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println("len of board: ", len(board))
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
