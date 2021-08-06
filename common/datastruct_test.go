package common

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	var l = ConstructList([]int{2, 1, 3})
	if fmt.Sprint(l) != "" {
		t.Errorf("Something looks wrong about the ListNode's String()")
	}
}
