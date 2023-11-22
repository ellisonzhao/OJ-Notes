package main

import (
	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = algorithm.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		// 左孩子节点全部入栈
		for curr != nil {
			stack = append(stack, curr)
			ans = append(ans, curr.Val)
			curr = curr.Left
		}
		// 节点出栈，尝试访问右孩子节点
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}
	return ans
}

func main() {

}
