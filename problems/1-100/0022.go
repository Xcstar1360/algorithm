package problem1_100

func generateParenthesis(n int) []string {
	var ans []string

	var dfs func(i, j, k int, s []byte)
	dfs = func(i, j, k int, s []byte) {
		if i == n && j == n && k == 0 {
			ans = append(ans, string(s))
			return
		}

		if i < n {
			dfs(i+1, j, k+1, append(s, '('))
		}
		if j < n && k > 0 {
			dfs(i, j+1, k-1, append(s, ')'))
		}
	}
	dfs(0, 0, 0, []byte{})

	return ans
}
