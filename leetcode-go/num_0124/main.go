package main

import (
	"math"

	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(dfs(node.Left), 0)
		rightSum := max(dfs(node.Right), 0)
		maxSum = max(leftSum+rightSum+node.Val, maxSum)
		return max(leftSum, rightSum) + node.Val
	}

	dfs(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
