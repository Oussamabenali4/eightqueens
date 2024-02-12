package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var board [8]int
	var solutions [][8]int
	y := 0
	y, board, solutions = solution(y, 1, board, solutions)
	solution_print(solutions)
}

func solution(y int, start int, board [8]int, solutions [][8]int) (int, [8]int, [][8]int) { // i should pass the function as paramiter .
	for i := start; i <= 8; i++ { // this is for the possible solutions
		if !diag(i, y, board) { // yo check if the solution is valid
			board[y] = i
			if y < 7 {
				return solution(y+1, 1, board, solutions) // stil has work to do
			} else {
				solutions = append(solutions, board)
				if board[y-1] == 8 { // return to the previous
					if y-1 == 0 {
						return y, [8]int{}, solutions
					}
					return solution(y-2, board[y-2]+1, board, solutions)
				}
				return solution(y-1, board[y-1]+1, board, solutions)
				//return y, board, solutions // done with the solution // here i need to find a way to move to the next
			}
		} else {
			if i < 8 { // if didn't find any solution just move to the next possible
				continue
			} else { // but if there is no solution at all
				if board[0] == 8 && board[1] == 6 && board[2] == 4 { // you just stop if there will be no solutions and return the signal
					return 0, [8]int{}, solutions
				}
				if board[y-1] == 8 { // return to the previous
					if y-1 == 0 {
						return y, [8]int{}, solutions
					}
					return solution(y-2, board[y-2]+1, board, solutions)
				}
				return solution(y-1, board[y-1]+1, board, solutions)
			}
		}
	}
	return y, board, solutions
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

func solution_print(solutions [][8]int) {
	for _, item := range solutions {
		for _, dig := range item {
			z01.PrintRune(rune(dig) + '0')
		}
		z01.PrintRune('\n')
	}
}
