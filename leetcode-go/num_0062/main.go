package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	sums := make([]int, n)
	for i := range sums {
		sums[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sums[j] += sums[j-1]
		}
	}

	return sums[n-1]
}
