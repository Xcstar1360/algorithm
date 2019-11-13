package linkedlist

/**
 * 判断是否为回文串
 * 数据结构 (SingleLinkedList with head ptr)
 */
func isPalindrome(l *LinkedList) bool {
	// 尽管后面代码能处理这两种边界条件,但是为了提升代码执行效率,直接返回
	if l.length == 0 || l.length == 1 {
		return true
	}

	var prev *ListNode
	// 带头指针的单链,跳过头指针
	var fast, slow = l.head.next, l.head.next
	for fast != nil && fast.next != nil {
		fast = fast.next.next

		// 链表反转
		var tmp = slow.next
		slow.next = prev
		prev = slow
		slow = tmp
	}

	// 复原链表使用
	var mid, left = slow, prev
	// 奇数个,跳过中间值
	if fast != nil {
		slow = slow.next
	}

	var result = true
	for slow != nil {
		if slow.value != prev.value {
			result = false
		}
		slow = slow.next
		prev = prev.next
	}

	// 复原链表
	for left != nil {
		var tmp = left.next
		left.next = mid
		mid = left
		left = tmp
	}
	l.head.next = mid

	return result
}
