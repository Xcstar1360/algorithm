package common

import "testing"

func TestConstructList(t *testing.T) {
	var nums = []int{3, 1, 4, 5, 2, 9, 10}
	var l = ConstructList(nums)
	for i, v := range nums {
		if v != l.Val {
			t.Errorf(
				"The number of index`%d` should be %d, but got %v",
				i,
				v,
				l.Val,
			)
		}
		l = l.Next
	}
}
