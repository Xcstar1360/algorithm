package one

import (
	c "github.com/Xcstar1360/algorithm/common"
)

func addTwoNumbers(l1 *c.ListNode, l2 *c.ListNode) *c.ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	var l = &c.ListNode{}
	var carry, p = 0, l
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		v := carry % 10
		carry /= 10

		l.Next = &c.ListNode{Val: v}
		l = l.Next
	}

	return p.Next
}
