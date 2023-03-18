package num_0695

// func maxAreaOfIsland(grid [][]int) int {
// 	if len(grid) == 0 {
// 		return 0
// 	}
//
// 	var ans int
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			ans = int(math.Max(float64(ans), float64(dfs(grid, i, j))))
// 		}
// 	}
// 	return ans
// }
//
// // 深度优先搜索
// func dfs(grid [][]int, i, j int) int {
// 	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
// 		return 0
// 	}
// 	// 标记当前位置已经访问过
// 	grid[i][j] = -1
// 	dx := []int{0, 1, 0, -1}
// 	dy := []int{1, 0, -1, 0}
// 	var area int
// 	for index := 0; index < 4; index++ {
// 		x, y := i+dx[index], j+dy[index]
// 		area += dfs(grid, x, y)
// 	}
// 	// 恢复当前位置
// 	grid[i][j] = 1
//
// 	return area
// }

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var ans int
	// 广度优先搜索
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			var area int
			queue := [][]int{{i, j}}
			for len(queue) > 0 {
				// 队头出队
				pairs := queue[0]
				queue = queue[1:]
				pi, pj := pairs[0], pairs[1]
				if pi < 0 || pi >= len(grid) || pj < 0 || pj >= len(grid[0]) || grid[pi][pj] != 1 {
					continue
				}
				// 标记已访问过
				grid[pi][pj] = -1
				area += 1
				// 遍历下一个位置
				dx := []int{0, 1, 0, -1}
				dy := []int{1, 0, -1, 0}
				for index := 0; index < 4; index++ {
					x, y := pi+dx[index], pj+dy[index]
					queue = append(queue, []int{x, y})
				}
			}
			ans = maxInt(area, ans)
		}
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
