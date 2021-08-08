package problem1101_1200

import "testing"

/**
 * 0 <= n <= 37
 * 答案确保是 32 位以内的整数
 * The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
 */

type testCase1137 struct {
	n   int
	ans int
}

var test_Cases1137 = []testCase1137{
	{n: 4, ans: 4},
	{n: 25, ans: 1389537},
}

func TestTribonacci(t *testing.T) {
	for _, v := range test_Cases1137 {
		var res = tribonacci(v.n)
		if res != v.ans {
			t.Errorf("The `%dth` number of tribonacci should be %d, but got %d", v.n+1, v.ans, res)
		}
	}
}
