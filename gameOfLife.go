package main

import "fmt"

func createBoardGame(rows, cols int) [][]int {
	board := make([][]int, rows)
	for indexRow := 0; indexRow < rows; indexRow++ {
		board[indexRow] = make([]int, cols)

	}
	return board

}

func main() {
	fmt.Println(createBoardGame(10, 10))
}
