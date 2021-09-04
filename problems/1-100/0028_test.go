package problem1_100

import "testing"

/**
 * 0 <= haystack.length, needle.length <= 5 * 10^4
 * haystack 和 needle 仅由小写英文字符组成
 */

type testCase28 struct {
	haystack string
	needle   string
	ans      int
}

var test_Cases28 = []testCase28{
	{
		haystack: "hello",
		needle:   "ll",
		ans:      2,
	},
	{
		haystack: "aaaaa",
		needle:   "bba",
		ans:      -1,
	},
	{
		haystack: "",
		needle:   "",
		ans:      0,
	},
}

func TestStrStrBF(t *testing.T) {
	for _, v := range test_Cases28 {
		if strStrBF(v.haystack, v.needle) != v.ans {
			t.Errorf("Answer Error")
		}
	}
}

func TestStrStrRK(t *testing.T) {
	for _, v := range test_Cases28 {
		if strStrRK(v.haystack, v.needle) != v.ans {
			t.Errorf("Answer Error")
		}
	}
}
