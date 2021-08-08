package problem1_100

import "testing"

/**
 * 1 <= s.length <= 1000
 * 字符串由英文字母(包含大小写)、',' 以及 '.' 组成
 * s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 1 <= numRows <= 1000
 */

type testCase6 struct {
	s   string
	row int
	ans string
}

var test_Cases6 = []testCase6{
	{s: "PAYPALISHIRING", row: 3, ans: "PAHNAPLSIIGYIR"},
	{s: "PAYPALISHIRING", row: 4, ans: "PINALSIGYAHRPI"},
	{s: "A", row: 1, ans: "A"},
}

func TestConvert(t *testing.T) {
	for _, v := range test_Cases6 {
		var res = convert(v.s, v.row)
		if res != v.ans {
			t.Errorf("\nThe result should be %s,\nbut got %s", v.ans, res)
		}
	}
}
