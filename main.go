package main

import (
	chess "eightqueens/packages"
	"fmt"
)

func PrintBoard(board [8][8]int) {
	for i := 0; i < 8; i++ {
		fmt.Println(board[i])
	}
}

type square struct {
	x int
	y int
}

func free(pos square, seen []square) bool {
	for i := 0; i < len(seen); i++ {
		if pos == seen[i] {
			return false
		}
	}
	return true
}

func onesolution(board [8][8]int, sl []square) ([8][8]int, []square) {
	var sq square
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			sq = square{x: i, y: j}
			if board[i][j] != 8 && board[i][j] != 1 && free(sq, sl) {
				// here i need a data type
				sl = append(sl, sq)
				board = chess.Queen(i+1, j+1, board)
				return onesolution(board, sl)
			} else {
				sl = append(sl, sq)
			}
		}
	}
	return board, sl
}

// put 8 queens on the board
func main() {
	var sl []square
	var board [8][8]int
	// trying one solutionasga
	board, sl = onesolution(board, sl)
	PrintBoard(board)
}