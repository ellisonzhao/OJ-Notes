package num_0994

func orangesRotting(grid [][]int) int {
	if grid == nil || grid[0] == nil {
		return 0
	}
	var queue [][]int
	var freshCount int
	// 统计初始烂的橘子个数
	row, col := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	// 开始 bfs 遍历
	var minutes int
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		minutes++
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x, y := point[0]+dir[0], point[1]+dir[1]
				if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 || grid[x][y] == 2 {
					continue
				}
				// 周围橘子腐烂
				grid[x][y] = 2
				queue = append(queue, []int{x, y})
				freshCount--
			}
		}
	}

	if freshCount > 0 {
		return -1
	}

	return minutes - 1
}
