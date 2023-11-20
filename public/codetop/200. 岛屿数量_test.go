package codetop

func renderLand(grid [][]byte, i, j, M, N int, index [][]bool) {
	if i < M && j < N && i >= 0 && j >= 0 && grid[i][j] == '1' && !index[i][j] {
		index[i][j] = true
		renderLand(grid, i-1, j, M, N, index)
		renderLand(grid, i+1, j, M, N, index)
		renderLand(grid, i, j+1, M, N, index)
		renderLand(grid, i, j-1, M, N, index)
	}
}

func numIslands(grid [][]byte) int {
	cnt := 0
	M := len(grid)
	N := len(grid[0])
	index := make([][]bool, M)
	for i := range index {
		index[i] = make([]bool, N)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' && !index[i][j] {
				renderLand(grid, i, j, M, N, index)
				cnt = cnt + 1
			}
		}
	}

	return cnt
}
