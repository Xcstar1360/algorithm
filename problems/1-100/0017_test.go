package problem1_100

/**
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 之间的一个数字
 * digits[i] is a digit in the range ['2', '9']
 * 因为本题不限制返回顺序,做验证有点麻烦.测试用例中只认为标准顺序才是正确答案
 */

type testCase17 struct {
	digits string
	ans    []string
}

var test_Cases17 = []testCase17{
	{
		digits: "23",
		ans: []string{
			"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
		},
	},
	{
		digits: "",
		ans:    []string{},
	},
	{
		digits: "2",
		ans:    []string{"a", "b", "c"},
	},
}

// func TestLetterCombinations(t *testing.T) {
// 	for _, v := range test_Cases17 {
// 		var res =
// 	}
// }
