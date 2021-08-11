package problem1_100

import (
	"testing"

	c "github.com/Xcstar1360/algorithm/common"
)

/**
 * 链表中的节点数为sz
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 */

type testCase19 struct {
	origin *c.ListNode
	head   *c.ListNode
	n      int
	ans    *c.ListNode
}

var test_Cases19 = []testCase19{
	{
		origin: c.ConstructList([]int{1, 2, 3, 4, 5}),
		head:   c.ConstructList([]int{1, 2, 3, 4, 5}),
		n:      2,
		ans:    c.ConstructList([]int{1, 2, 3, 5}),
	},
	{
		origin: c.ConstructList([]int{1}),
		head:   c.ConstructList([]int{1}),
		n:      1,
		ans:    nil,
	},
	{
		origin: c.ConstructList([]int{1, 2}),
		head:   c.ConstructList([]int{1, 2}),
		n:      1,
		ans:    c.ConstructList([]int{1}),
	},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, v := range test_Cases19 {
		var res = removeNthFromEnd(v.head, v.n)
		if !c.ListEq(res, v.ans) {
			t.Errorf(
				"\nRemove the %dth node from the end of the list:\n%v,\nthe result should be:\n%v,\nbut got:\n%v",
				v.n,
				v.origin,
				v.ans,
				res,
			)
		}
	}
}
