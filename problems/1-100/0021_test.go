package problem1_100

import (
	"testing"

	c "github.com/Xcstar1360/algorithm/common"
)

/**
 * 每条链表的节点数在 [0, 50] 之间
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * l1 与 l2 均为非递减排列
 * Both l1 and l2 are sorted in non-decreasing order.
 */

type testCase21 struct {
	l1  []int
	l2  []int
	ans *c.ListNode
}

var test_Cases21 = []testCase21{
	{
		l1:  []int{1, 2, 4},
		l2:  []int{1, 3, 4},
		ans: c.ConstructList([]int{1, 1, 2, 3, 4, 4}),
	},
	{
		l1:  []int{},
		l2:  []int{},
		ans: nil,
	},
	{
		l1:  []int{},
		l2:  []int{0},
		ans: c.ConstructList([]int{0}),
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, v := range test_Cases21 {
		var res = mergeTwoLists(
			c.ConstructList(v.l1),
			c.ConstructList(v.l2),
		)
		if !c.ListEq(res, v.ans) {
			t.Errorf(
				"\nMerge two sorted list:\n%v and %v\nthe result should be %v\nbut got %v",
				v.l1,
				v.l2,
				v.ans,
				res,
			)
		}
	}
}
