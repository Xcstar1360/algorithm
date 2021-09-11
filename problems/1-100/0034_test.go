package problem1_100

import (
	"testing"

	c "github.com/Xcstar1360/algorithm/common"
)

/**
 * 0 <= nums.length <= 105
 * -109 <= nums[i] <= 109
 * nums 是一个非递减数组
 * -109 <= target <= 109
 */

type testCase34 struct {
	nums   []int
	target int
	ans    []int
}

var test_Cases34 = []testCase34{
	{
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 8,
		ans:    []int{3, 4},
	},
	{
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 6,
		ans:    []int{-1, -1},
	},
	{
		nums:   []int{},
		target: 0,
		ans:    []int{-1, -1},
	},
	{
		nums:   []int{1},
		target: 1,
		ans:    []int{0, 0},
	},
}

func TestSearchRange(t *testing.T) {
	for _, v := range test_Cases34 {
		var res = searchRange(v.nums, v.target)
		if !c.SliceEq(res, v.ans) {
			t.Error("Answer Error")
		}
	}
}
