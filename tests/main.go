package main

import "fmt"

func main() {
	var sl []int
	var board [8]int
	board = solution(0, board, sl)
	fmt.Println(board)
}

// the enitial value of y will be 0 and will be increasing

func solution(y int, board [8]int, sl []int) [8]int {
	for i := 1; i <= 8; i++ {
		if !in(i, sl) { //&& i != sl[y-1] && i != sl[y-1]-1 {
			if !diag(i, sl) {
				sl = append(sl, i)
				board[y] = i
				if y < 7 {
					return solution(y+1, board, sl)
				} else {
					return board
				}
			}
		}
	}
	return board
}
func in(i int, sl []int) bool {
	if len(sl) == 0 {
		return false
	}
	for _, item := range sl {
		if item == i {
			return true
		}
	}
	return false
}

func diag(i int, sl []int) bool { // this needs further devlopment; 
	if len(sl) == 0 {
		return false
	}
	for j, item := range sl {
		if item == i+j+1 || item == i-j-1 {
			return true
		}
	}
	return false
}
