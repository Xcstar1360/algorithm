package problem1_100

import "testing"

/**
 * 0 <= nums.length <= 3 * 104
 * -100 <= nums[i] <= 100
 * nums 已经按升序排列
 * nums is sorted in non-decreasing order.
 */

type testCase26 struct {
	nums []int
	ans  int
}

var test_Cases26 = []testCase26{
	{
		nums: []int{1, 1, 2},
		ans:  2,
	},
	{
		nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		ans:  5,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, v := range test_Cases26 {
		if v.ans != removeDuplicates(v.nums) {
			t.Error("Answer error")
		}
	}
}
