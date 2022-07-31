package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	sums := make([]int, n)
	for i := 0; i < n; i++ {
		sums[i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			sums[j] = min(triangle[i][j]+sums[j], triangle[i][j]+sums[j+1])
		}
	}

	return sums[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
