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
		//if !in(i, sl) { //&& i != sl[y-1] && i != sl[y-1]-1 {
		if !diag(i, y, board) {
			sl = append(sl, i)
			board[y] = i
			if y < 7 {
				return solution(y+1, board, sl)
			} else {
				return board
			}
		}
		//}
	}
	return board
}

// func in(i int, sl []int) bool {
// 	if len(sl) == 0 {
// 		return false
// 	}
// 	for _, item := range sl {
// 		if item == i {
// 			return false // true
// 		}
// 	}
// 	return false
// }

// i + y - j
func diag(i int, y int, board [8]int) bool { // this needs further devlopment;
	if board[0] == 0 {
		return false
	}
	for j := 0; j < y; j++ {
		if board[j] == 0 {
			return false
		}
		if board[j]+y-j == i || board[j]-y+j == i || board[j] == i { // i need to fix this equation
			return true
		}

	}
	return false
}
