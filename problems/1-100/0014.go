package problem1_100

func longestCommonPrefix(strs []string) string {
	var ans []byte

	var round = 0
	for {
		if len(strs[0]) <= round {
			break
		}
		var i, ch = 1, strs[0][round]
		for ; i < len(strs); i++ {
			if len(strs[i]) <= round || strs[i][round] != ch {
				break
			}
		}
		if i != len(strs) {
			break
		}
		round++
		ans = append(ans, ch)
	}

	return string(ans)
}
