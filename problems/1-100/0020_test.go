package problem1_100

import "testing"

/**
 * 1 <= s.length <= 104
 * s 仅由 '()[]{}' 字符组成
 * s consists of parentheses only '()[]{}'.
 */

type testCase20 struct {
	s   string
	ans bool
}

var test_Cases20 = []testCase20{
	{s: "()", ans: true},
	{s: "()[]{}", ans: true},
	{s: "(]", ans: false},
	{s: "([)]", ans: false},
	{s: "{[]}", ans: true},
	{s: "[", ans: false},
}

func TestIsValid(t *testing.T) {
	for _, v := range test_Cases20 {
		var res = isValid(v.s)
		if res != v.ans {
			if v.ans {
				t.Errorf("The string `%s` is valid, but got false", v.s)
			} else {
				t.Errorf("The string `%s` is not valid, but got true", v.s)
			}
		}
	}
}
