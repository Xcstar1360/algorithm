package problem1_100

import (
	c "github.com/Xcstar1360/algorithm/common"
)

func mergeTwoLists(l1 *c.ListNode, l2 *c.ListNode) *c.ListNode {
	var p = &c.ListNode{}
	var l = p
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return l.Next
}
