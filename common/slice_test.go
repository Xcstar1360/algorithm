package common

import (
	"testing"
)

func TestSliceEq(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 3, 2, 4}
	c := []int{1, 2, 3, 4}
	d := []int{1, 2, 3}

	if SliceEq(a, b) {
		t.Errorf("%v should not equals to %v", a, b)
	}
	if !SliceEq(a, c) {
		t.Errorf("%v should equals to %v", a, c)
	}
	if SliceEq(a, d) {
		t.Errorf("%v should not equals to %v", a, d)
	}
}
