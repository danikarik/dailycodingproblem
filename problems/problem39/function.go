package problem39

func playGameOfLife(board [][]int, n int) [][]string {
	for i := 0; i < n; i++ {
		board = nextGen(board)
	}
	return printBoard(board)
}

func printBoard(board [][]int) [][]string {
	n := len(board)
	m := len(board[0])
	display := make([][]string, n)
	for i := 0; i < n; i++ {
		display[i] = make([]string, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 1 {
				display[i][j] = "*"
				continue
			}
			display[i][j] = "."
		}
	}
	return display
}

func nextGen(board [][]int) [][]int {
	n := len(board)
	m := len(board[0])
	future := make([][]int, n)
	for i := 0; i < n; i++ {
		future[i] = make([]int, m)
	}
	// loop through every cell
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			// finding no of neighbours that are alive
			aliveNeighbours := 0
			for t := -1; t <= 1; t++ {
				for y := -1; y <= 1; y++ {
					aliveNeighbours += board[i+t][j+y]
				}
			}

			// the cell needs to be subtracted from
			// its neighbours as it was counted before
			aliveNeighbours -= board[i][j]

			// implementing the rules of game
			if board[i][j] == 1 && aliveNeighbours < 2 {
				// cell is lonely and dies
				future[i][j] = 0
			} else if board[i][j] == 1 && aliveNeighbours > 3 {
				// cell dies due to over population
				future[i][j] = 0
			} else if board[i][j] == 0 && aliveNeighbours == 3 {
				// a new cell is born
				future[i][j] = 1
			} else {
				// remains the same
				future[i][j] = board[i][j]
			}
		}
	}
	return future
}
