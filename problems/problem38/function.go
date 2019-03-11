package problem38

func countQueens(board [][]int) int {
	cnt := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func isSafe(board [][]int, row, col int) bool {
	// check this row on left side
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}
	// check upper diagonal on left side
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	// check lower diagonal on left side
	for i, j := row, col; j >= 0 && i < len(board); i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func makeBoard(n int) [][]int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	return board
}

func placeQueens(board [][]int, col int) bool {
	n := len(board)
	// base case: if all queens are placed then return true
	if col >= n {
		return true
	}
	// consider this column and try placing this queen in all rows one by one
	for i := 0; i < n; i++ {
		// check if the queen can be placed on board[i][col]
		if isSafe(board, i, col) {
			// place this queen in board[i][col]
			board[i][col] = 1
			// recur to place rest of the queens
			if placeQueens(board, col+1) {
				return true
			}
			board[i][col] = 0
		}
	}
	return false
}

func queens(n int) int {
	board := makeBoard(n)
	if !placeQueens(board, 0) {
		return 0
	}
	return countQueens(board)
}
