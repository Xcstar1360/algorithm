package problem1_100

import (
	c "github.com/Xcstar1360/algorithm/common"
)

func removeNthFromEnd(head *c.ListNode, n int) *c.ListNode {
	var h = &c.ListNode{Next: head}
	var p, l = h, h
	for n > 0 {
		l = l.Next
		n--
	}
	for l.Next != nil {
		p = p.Next
		l = l.Next
	}
	p.Next = p.Next.Next

	return h.Next
}
