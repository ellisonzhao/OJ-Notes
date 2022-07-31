package main

func main() {

}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	sums := make([]int, n)

	sums[0] = grid[0][0]
	for i := 1; i < n; i++ {
		sums[i] = grid[0][i] + sums[i-1]
	}

	for i := 1; i < m; i++ {
		sums[0] = sums[0] + grid[i][0]
		for j := 1; j < n; j++ {
			sums[j] = min(sums[j-1]+grid[i][j], sums[j]+grid[i][j-1])
		}
	}

	return sums[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
