package main

import "fmt"

func PrintBoard(board [8][8]int) {
	for i := 0; i < 8; i++ {
		fmt.Println(board[i])
	}
}

// position and controled squares by one queen
func Queen(y int, x int, board [8][8]int) [8][8]int {
	x = x - 1
	y = y - 1

	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			// y = x + cst ; cst = y - x
			if i == j+x-y {
				board[j][i] = 1
			}
			// y = -x + b ; cst = y + x
			if i == -j+(y+x) {
				board[j][i] = 1
			}
			if j == x {
				board[y][i] = 1
			}
			if i == y {
				board[j][x] = 1
			}
		}
	}
	// the position of the queen gits the number 7
	board[y][x] = 7
	return board
}

func main() {
	var board [8][8]int
	// value for testiig 3 4
	board = Queen(5, 3, board)
	PrintBoard(board)
}
