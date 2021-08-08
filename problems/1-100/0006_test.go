package problem1_100

import "testing"

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
