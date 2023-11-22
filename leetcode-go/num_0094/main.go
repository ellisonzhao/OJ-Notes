package main

import (
	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	curr := root
	var stack []*TreeNode
	for curr != nil || len(stack) > 0 {
		// 左孩子节点入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 从最左孩子节点开始遍历
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		// 访问右子树
		curr = curr.Right
	}
	return ans
}

func main() {

}
