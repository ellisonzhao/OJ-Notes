package main

import (
	"strconv"
	"strings"

	"github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"
)

func main() {

}

type TreeNode = algorithm.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var values []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			values = append(values, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			values = append(values, "null")
		}
	}
	return strings.Join(values, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	values := strings.Split(data, ",")
	if len(values) <= 0 {
		return nil
	}

	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	idx := 1
	for idx < len(values) {
		node := q[0]
		q = q[1:]
		leftVal := values[idx]
		rightVal := values[idx+1]
		if leftVal != "null" {
			v, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: v}
			node.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "null" {
			v, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: v}
			node.Right = rightNode
			q = append(q, rightNode)
		}
		idx += 2
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
