package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	tail := &ListNode{
		Val:  2,
		Next: nil,
	}

	n3 := &ListNode{
		Val:  9,
		Next: tail,
	}

	n2 := &ListNode{
		Val:  3,
		Next: n3,
	}

	n1 := &ListNode{
		Val:  4,
		Next: n2,
	}

	head := &ListNode{
		Val:  1,
		Next: n1,
	}

	head = sortList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 计算链表长度
	length := 1
	for node := head; node.Next != nil; node = node.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	for step := 1; step < length; step <<= 1 {
		prev, curr := dummy, dummy.Next
		// 按步长对子链表排序
		for curr != nil {
			left := curr
			right := splitListWithStep(left, step)
			curr = splitListWithStep(right, step)

			prev.Next = mergeSortedList(left, right)
			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}
	return dummy.Next
}

func splitListWithStep(head *ListNode, step int) *ListNode {
	if head == nil {
		return head
	}

	for i := 1; i < step && head.Next != nil; i++ {
		head = head.Next
	}

	right := head.Next
	head.Next = nil

	return right
}

func mergeSortedList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}
