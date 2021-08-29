package problem1_100

import "testing"

/**
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 50
 * 0 <= val <= 100
 */

type testCase27 struct {
	nums []int
	val  int
	ans  int
}

var test_Cases27 = []testCase27{
	{
		nums: []int{2},
		val:  3,
		ans:  1,
	},
	{
		nums: []int{3, 2, 2, 3},
		val:  3,
		ans:  2,
	},
	{
		nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:  2,
		ans:  5,
	},
}

func TestRemoveElement(t *testing.T) {
	for _, v := range test_Cases27 {
		if v.ans != removeElement(v.nums, v.val) {
			t.Error("Answer error")
		}
	}
}
