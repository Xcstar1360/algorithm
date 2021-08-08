package problem1_100

import "testing"

/**
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * 字符串数组中的每个字符串都仅包含小写字母
 * strs[i] consists of only lower-case English letters.
 */

type testCase14 struct {
	strs []string
	ans  string
}

var test_Cases14 = []testCase14{
	{
		strs: []string{
			"flower", "flow", "flight",
		},
		ans: "fl",
	},
	{
		strs: []string{
			"dog", "racecar", "car",
		},
		ans: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range test_Cases14 {
		var res = longestCommonPrefix(v.strs)
		if res != v.ans {
			t.Errorf(
				"The longest common prefix string amongst `%v` should be \"%s\", but got \"%s\"",
				v.strs,
				v.ans,
				res,
			)
		}
	}
}
