package problem1_100

import "testing"

/**
 * 1 <= nums.length <= 104
 * -104 <= nums[i] <= 104
 * nums 为无重复元素的升序排列数组
 * -104 <= target <= 104
 */

type testCase35 struct {
	nums   []int
	target int
	ans    int
}

var test_Cases35 = []testCase35{
	{
		nums:   []int{1, 3, 5, 6},
		target: 5,
		ans:    2,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 2,
		ans:    1,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 7,
		ans:    4,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 0,
		ans:    0,
	},
	{
		nums:   []int{1},
		target: 0,
		ans:    0,
	},
}

func TestSearchInsert(t *testing.T) {
	for _, v := range test_Cases35 {
		if searchInsert(v.nums, v.target) != v.ans {
			t.Error("Answer Error")
		}
	}
}
