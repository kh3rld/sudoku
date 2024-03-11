package main

import (
	"fmt"
	"os"
	"sudoku"
)

func main() {
	args := os.Args
	if !sudoku.CheckParam(args[1:]) || !sudoku.Solution(args[1:]) {
		fmt.Println("Error")
	}
}
