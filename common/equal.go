package common

/**
 * 测试两个切片每个位置的元素是否相等
 * 因为Go暂时不支持泛型,所以先使用int
 */
func SliceEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

/**
 * 测试两条链表是否相等(这里只比较每个节点值,不考虑相交或部分重合的情况)
 * 因为Go暂时不支持泛型,所以先以int为例
 */
func ListEq(l1, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			return false
		}
		if l1 != nil && l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
