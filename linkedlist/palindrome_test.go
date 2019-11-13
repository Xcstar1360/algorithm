package linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	array := map[string]bool{
		"":          true,
		"1":         true,
		"12321":     true,
		"123321":	 true,
		" ":         true,
		"上海自来水来自海上": true,
		"abcde":     false,
	}
	for str, exp := range array {
		var l = NewLinkedList()
		for _, ch := range str {
			l.InsertToTail(ch)
		}
		var res = isPalindrome(l)
		if exp != res {
			t.Errorf("%v expected: %t, but got: %t\n", str, exp, res)
		}
	}
}
