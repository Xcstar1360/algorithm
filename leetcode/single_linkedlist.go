package main

/**
 * 单链表题目
 * leetcode@19	删除单链表的倒数第N个节点
 * leetcode@21	合并两个有序链表
 * leetcode@141 单链环形检测
 * leetcode@160 相交链表
 * leetcode@206 单链表反转
 * leetcode@876 单链中间节点
 */

/**
 * 单链节点
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 带头指针的单链
 */
type LinkedList struct {
	head *ListNode
}

/**
 * 单链反转
 */
func reverseList(head *ListNode) *ListNode {
	var tmp, rev *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = rev
		rev = head
		head = tmp
	}
	return rev
}

/**
 * 单链是否有环
 */
func hasCycle(head *ListNode) bool {
	if head != nil {
		var fast, slow = head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next

			if fast == slow {
				return true
			}
		}
	}

	return false
}

/**
 * 合并两个有序单链表
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 以 l1 为基础
	var cur, prev *ListNode
	for l2 != nil {
		cur = l1

		for cur != nil && l2.Val > cur.Val {
			prev = cur
			cur = cur.Next
		}
		if cur == l1 {
			var newNode = &ListNode{l2.Val, nil}
			newNode.Next = l1
			l1 = newNode
		} else if cur == nil {
			prev.Next = l2
			break
		} else {
			var newNode = &ListNode{l2.Val, nil}
			newNode.Next = cur
			prev.Next = newNode
		}

		l2 = l2.Next
	}

	return l1
}

/**
 * 删除单链表的倒数第 n 个节点
 * 题目保证 n 一定有效,这里不做判断
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var front, back = head, head
	for i := 0; i < n; i++ {
		back = back.Next
	}

	if back == nil {
		// 删除 head 节点
		head = head.Next
	} else {
		// 仔细品味这里用 back.Next 而不用 back
		for back.Next != nil {
			back = back.Next
			front = front.Next
		}
		front.Next = front.Next.Next
	}

	return head
}

/**
 * 返回单链中间节点
 * 如果有两个, 返回第二个
 */
func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/**
 * 相交链表
 * 他人解法: 不是最快,但是代码最简单
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var pA, pB = headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
