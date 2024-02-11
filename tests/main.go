package main

import "fmt"

func main() {
	var board [8]int
	board = solution(0, 1, board)
	fmt.Println(board)
}

func solution(y int, start int, board [8]int) [8]int {
	for i := start; i <= 8; i++ {
		if !diag(i, y, board) {
			board[y] = i
			if y < 7 {
				return solution(y+1, 1, board)
			} else {
				return board
			}
		} else {
			if i < 8 {
				continue
			} else {
				if board[y-1] == 8 {
					return solution(y-2, board[y-2]+1, board)
				}
				return solution(y-1, board[y-1]+1, board)
			}
		}
	}
	return board
}

func diag(i int, y int, board [8]int) bool {
	if board[0] == 0 {
		return false
	}
	for j := 0; j < y; j++ {
		if board[j] == 0 {
			return false
		}
		if board[j]+y-j == i || board[j]-y+j == i || board[j] == i {
			return true
		}

	}
	return false
}
