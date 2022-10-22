package num_0200

func numIslands(grid [][]byte) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, row, col int) int {
	// 边界
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return 0
	}

	// 非岛屿
	if grid[row][col] != '1' {
		return 0
	}

	// 标记该岛屿已经访问过(每一个点只需要访问一次)
	grid[row][col] = 0

	// 继续遍历剩余岛屿
	return 1 +
		dfs(grid, row, col+1) +
		dfs(grid, row+1, col+1) +
		dfs(grid, row, col-1) +
		dfs(grid, row-1, col)
}
