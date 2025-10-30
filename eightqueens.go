package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := make([]int, 8)
	solve(board, 0)
}

func solve(board []int, col int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row + 1
			solve(board, col+1)
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row+1 ||
			board[i] == row+1-(col-i) ||
			board[i] == row+1+(col-i) {
			return false
		}
	}
	return true
}

func printBoard(board []int) {
	for _, pos := range board {
		z01.PrintRune(rune('0' + pos))
	}
	z01.PrintRune('\n')
}
