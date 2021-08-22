package problem1_100

import (
	c "github.com/Xcstar1360/algorithm/common"
)

func swapPairs(head *c.ListNode) *c.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var l = &c.ListNode{Next: head}
	var p = head
	head = l
	for p != nil && p.Next != nil {
		l.Next = p.Next
		// 不推荐这种写法,但是Golang这么写很方便
		p.Next, p.Next.Next = p.Next.Next, p

		l = l.Next.Next
		p = p.Next
	}

	return head.Next
}
