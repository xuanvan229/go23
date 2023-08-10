package main

func countRectangles(board [][]int) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}

	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 1 && !visited[i][j] {
				count++
				dfs(board, i, j, visited)
			}
		}
	}
	return count
}

func dfs(board [][]int, row, col int, visited [][]bool) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == 0 || visited[row][col] {
		return
	}

	// Mark the cell as visited
	visited[row][col] = true

	// Explore horizontally
	for i := col + 1; i < len(board[0]); i++ {
		if board[row][i] == 0 || visited[row][i] {
			break
		}
		dfs(board, row, i, visited)
	}

	// Explore vertically
	for i := row + 1; i < len(board); i++ {
		if board[i][col] == 0 || visited[i][col] {
			break
		}
		dfs(board, i, col, visited)
	}
}
