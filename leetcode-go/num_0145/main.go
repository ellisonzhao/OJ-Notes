package main

import (
	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	curr := root
	var prev *TreeNode
	var stack []*TreeNode
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 只有当栈顶节点的左右孩子全部访问完才可以出栈
		node := stack[len(stack)-1]
		if node.Right != nil && node.Right != prev {
			// 右孩子节点还未访问
			curr = node.Right
		} else {
			// 出栈
			ans = append(ans, node.Val)
			prev = node
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

func main() {

}
