package problem1_100

import "testing"

/**
 * -2^31 <= x <= 2^31 - 1
 */

type testCase9 struct {
	x   int
	ans bool
}

var test_Cases9 = []testCase9{
	{x: 121, ans: true},
	{x: -121, ans: false},
	{x: 10, ans: false},
	{x: -101, ans: false},
}

func TestIsPalindrome(t *testing.T) {
	for _, v := range test_Cases9 {
		if isPalindrome(v.x) != v.ans {
			if v.ans {
				t.Errorf("`%d` is a palindrome number, but got false", v.x)
			} else {
				t.Errorf("`%d` is not a palindrome number, but got true", v.x)
			}
		}
	}
}
