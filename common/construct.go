package common

func ConstructList(nums []int) *ListNode {
	var p = &ListNode{}
	var l = p
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return l.Next
}
