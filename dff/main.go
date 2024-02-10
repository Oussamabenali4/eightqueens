package main

import "fmt"

func main() {

	var board [8]int
	var condition bool = true
	board, condition = solution(board, condition)
	fmt.Println(board)

}

func solution(board [8]int, condition bool) ([8]int, bool) { // this is suppose to give all the solutions. so it must return a slice
	for i := 1; i <= 8; i++ { // this is to loop over the board
		if board[i-1] == 0 { // if there is an empty place in it
			// find a solution//
			add := possible(board)
			if add == 0 {
				return board, false
			} else {
				board[i-1] = add
				return solution(board, true)
			}
			// if there is no solution return to 0 and and remmember to not go in the same path
			// if ther is no solution for the first square , stop all the madness
		}
	}
	return board, true // i stil have to figure out something here
}

func possible(board [8]int) int {
	for i := 1; i <= 8; i++ {
		// if the number is not in the board add it
		if board[i-1] == 0 {
			return i
		}
	}
	return 0
}
