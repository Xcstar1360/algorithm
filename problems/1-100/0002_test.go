package problem1_100

import (
	"testing"

	c "github.com/Xcstar1360/algorithm/common"
)

/**
 * 每条链表上的节点数在 1~100 之间
 * The number of nodes in each linked list is in the range [1, 100].
 * 0 <= Node.val <= 9
 * 保证链表表示的数字没有前导0
 * It is guaranteed that the list represents a number that does not have leading zeros.
 */

type testCase2 struct {
	l1  *c.ListNode
	l2  *c.ListNode
	ans *c.ListNode
}

var test_Cases2 = []testCase2{
	{
		l1:  c.ConstructList([]int{2, 4, 3}),
		l2:  c.ConstructList([]int{5, 6, 4}),
		ans: c.ConstructList([]int{7, 0, 8}),
	},
	{
		l1:  c.ConstructList([]int{0}),
		l2:  c.ConstructList([]int{0}),
		ans: c.ConstructList([]int{0}),
	},
	{
		l1:  c.ConstructList([]int{9, 9, 9, 9, 9, 9, 9}),
		l2:  c.ConstructList([]int{9, 9, 9, 9}),
		ans: c.ConstructList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, v := range test_Cases2 {
		var res = addTwoNumbers(v.l1, v.l2)
		if !c.ListEq(res, v.ans) {
			t.Errorf(
				"\nlist1: %v + list2: %v\nthe result should be %v\nbut got %v",
				v.l1,
				v.l2,
				v.ans,
				res,
			)
		}
	}
}
