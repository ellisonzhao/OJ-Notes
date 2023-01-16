package num_0142

import "github.com/ellisonzhao/OJ-Solution/leetcode-go/algorithm"

type ListNode = algorithm.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			pos := head
			for pos != slow {
				pos = pos.Next
				slow = slow.Next
			}

			return pos
		}
	}

	return nil
}
