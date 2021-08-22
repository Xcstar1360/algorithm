package problem1_100

import "testing"

/**
 * 1 <= n <= 8
 */

type testCase22 struct {
	n   int
	ans []string
}

var test_Cases22 = []testCase22{
	{
		n:   1,
		ans: []string{"()"},
	},
	{
		n:   2,
		ans: []string{"(())", "()()"},
	},
	{
		n: 3,
		ans: []string{
			"((()))", "(()())", "(())()", "()(())", "()()()",
		},
	},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, v := range test_Cases22 {
		var res = generateParenthesis(v.n)
		if len(v.ans) != len(res) {
			t.Error("Answer error")
		}
		for i := range v.ans {
			if v.ans[i] != res[i] {
				t.Error("Answer error")
			}
		}
	}
}
