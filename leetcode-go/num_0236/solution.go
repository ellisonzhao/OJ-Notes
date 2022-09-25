package num_0236

import (
	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}
