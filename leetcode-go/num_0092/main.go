package num_0092

import "github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"

type ListNode = algorithm.ListNode

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	dummy := &ListNode{Next: head}
// 	prevTail, subTail := dummy, head
// 	for i := 1; i < left; i++ {
// 		prevTail = subTail
// 		subTail = subTail.Next
// 	}
// 	for i := left + 1; i <= right; i++ {
// 		curr := subTail.Next
// 		subTail.Next = curr.Next
// 		curr.Next = prevTail.Next
// 		prevTail.Next = curr
// 	}
//
// 	return dummy.Next
// }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev, curr := dummy, head
	for i := 1; i < left; i++ {
		prev = curr
		curr = curr.Next
	}

	subTail := curr
	prev.Next = nil
	for i := left; i <= right; i++ {
		next := curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr = next
	}
	subTail.Next = curr

	return dummy.Next
}
