package problem1_100

import (
	"testing"

	"github.com/Xcstar1360/algorithm/common"
)

/**
 * 2 <= nums.length <= 104
 * -109 <= nums[i] <= 109
 * -109 <= target <= 109
 * 只会有一个正确答案存在
 * Only one valid answer exists
 */

type testCase1 struct {
	nums   []int
	target int
	ans    []int
}

var test_Cases1 = []testCase1{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		ans:    []int{0, 1},
	},
	{
		nums:   []int{3, 2, 4},
		target: 6,
		ans:    []int{1, 2},
	},
	{
		nums:   []int{3, 3},
		target: 6,
		ans:    []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, v := range test_Cases1 {
		var res = twoSum(v.nums, v.target)
		if !common.SliceEq(res, v.ans) {
			t.Errorf(
				"nums: %v, target: %d, the result should be %v, but got %v",
				v.nums,
				v.target,
				v.ans,
				res,
			)
		}
	}
}
