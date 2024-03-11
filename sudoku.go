package sudoku

import "github.com/01-edu/z01"

func CheckParam(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for i := range args {
		if len(args[i]) != 9 {
			return false
		}
		for j := range args[i] {
			l := args[i][j]
			if l != '.' && (l < '1' || l > '9') {
				return false
			}
		}
	}
	return true
}

func Board(board [][]byte) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if x > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(board[y][x]))
		}
		z01.PrintRune('\n')
	}
}

func checkSudoku(deep int, board [][]byte) bool {
	x := deep % 9
	y := deep / 9
	for i := 0; i < 9; i++ {
		if i != x && board[y][i] == board[y][x] {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if i != y && board[i][x] == board[y][x] {
			return false
		}
	}
	caseX := x / 3
	caseY := y / 3
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			xx := caseX*3 + i
			yy := caseY*3 + j
			if (xx != x || yy != y) && board[yy][xx] == board[y][x] {
				return false
			}
		}
	}
	return true
}

func solve(deep int, board [][]byte) bool {
	if deep == 81 {
		Board(board)
		return true
	}
	x := deep % 9
	y := deep / 9
	if board[y][x] != '.' {
		if checkSudoku(deep, board) && solve(deep+1, board) {
			return true
		}
	} else {
		for i := '1'; i <= '9'; i++ {
			board[y][x] = byte(i)
			if checkSudoku(deep, board) && solve(deep+1, board) {
				return true
			}
		}
		board[y][x] = '.'
	}
	return false
}

func Solution(args []string) bool {
	var board [][]byte
	for i := range args {
		board = append(board, []byte(args[i]))
	}
	return solve(0, board)
}
