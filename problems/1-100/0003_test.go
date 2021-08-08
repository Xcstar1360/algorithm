package problem1_100

import "testing"

type testCase3 struct {
	s   string
	ans int
}

var test_Cases3 = []testCase3{
	{s: "abcabcbb", ans: 3},
	{s: "bbbbb", ans: 1},
	{s: "pwwkew", ans: 3},
	{s: "abba", ans: 2},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, v := range test_Cases3 {
		res := lengthOfLongestSubstring(v.s)
		if res != v.ans {
			t.Errorf(
				"\nAbout `%s`, the length of the longest substring without repeating charatcters should be %d, but got %d",
				v.s,
				v.ans,
				res,
			)
		}
	}
}
