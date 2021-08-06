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
