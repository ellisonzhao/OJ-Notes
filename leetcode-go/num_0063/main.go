package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 1; i < m; i++ {
		if dp[i][0] = dp[i-1][0]; obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		}
	}

	for j := 1; j < n; j++ {
		if dp[0][j] = dp[0][j-1]; obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] = dp[i-1][j] + dp[i][j-1]; obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}
