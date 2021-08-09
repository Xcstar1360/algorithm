package problem1_100

var mapping = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}

	var dfs func(count int, str []byte)
	dfs = func(count int, str []byte) {
		if count == len(digits) {
			ans = append(ans, string(str))
			return
		}
		var arr, _ = mapping[digits[count]]
		for _, ch := range arr {
			dfs(count+1, append(str, ch))
		}
	}
	dfs(0, []byte{})

	return ans
}
