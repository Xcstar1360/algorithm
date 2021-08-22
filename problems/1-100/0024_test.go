package problem1_100

import (
	"testing"

	c "github.com/Xcstar1360/algorithm/common"
)

/**
 * 链表中节点数在范围 [0, 100] 内
 * The number of nodes in linked list is in the range [0, 100].
 */

type testCase24 struct {
	head *c.ListNode
	ans  *c.ListNode
}

var test_Cases24 = []testCase24{
	{
		head: c.ConstructList([]int{1, 2, 3, 4}),
		ans:  c.ConstructList([]int{2, 1, 4, 3}),
	},
	{
		head: nil,
		ans:  nil,
	},
	{
		head: c.ConstructList([]int{1}),
		ans:  c.ConstructList([]int{1}),
	},
}

func TestSwapPairs(t *testing.T) {
	for _, v := range test_Cases24 {
		var res = swapPairs(v.head)
		if !c.ListEq(v.ans, res) {
			t.Errorf(
				"\nThe result should be\n\t%v\nbut got\n\t%v",
				v.ans,
				res,
			)
		}
	}
}
