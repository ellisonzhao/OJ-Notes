package main

import (
	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}

		dfs(node.Right)
	}

	dfs(root)

	return ans
}
