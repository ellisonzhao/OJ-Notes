package main

import (
	"fmt"

	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

type TreeNode = algorithm.TreeNode

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var isValid func(node *TreeNode) bool
	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// 递归访问左子树
		isLeftValid := isValid(node.Left)
		fmt.Println(isLeftValid)
		// 左子树节点值必须小于根结点
		if prev != nil && prev.Val >= node.Val {
			return false
		}
		// 更新当前访问节点
		prev = node
		isRightValid := isValid(node.Right)

		return isLeftValid && isRightValid
	}

	return isValid(root)
}

func main() {

}
