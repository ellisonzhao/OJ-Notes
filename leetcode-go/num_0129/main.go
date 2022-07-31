package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node5 := &TreeNode{5, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node9 := &TreeNode{9, node5, node1}
	node0 := &TreeNode{0, nil, nil}
	root := &TreeNode{4, node9, node0}

	fmt.Println(sumNumbers(root))
}

var res int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return res
	}

	preorder(root, 0)

	return res
}

func preorder(root *TreeNode, num int) {
	if root == nil {
		return
	}

	num += root.Val
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		res += num
		return
	}

	preorder(root.Left, num*10)
	preorder(root.Right, num*10)
}
