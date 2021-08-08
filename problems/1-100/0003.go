package problem1_100

/**
 * 这题解法容易想到,但不容易写好.尤其是左边界并不是碰到相同字符就一定滑动的
 * 思路:
 * 遍历时的下标 i 作为右边界,定义 left 为左边界,每一轮循环遍历到的字符为 ch
 * 当 ch 之前没有出现过,此时的 i 就是新的右边界
 * 当 ch 之前已经出现,从 hash 中取出它之前出现的位置 idx,当左边界小于等于 idx 时才滑动到 idx 下一位
 * 最后比较 i-left+1(即最新的不重复子串的长度) 与 maxLen 的值,较大值更新到 maxLen
 */
func lengthOfLongestSubstring(s string) int {
	var hash = make(map[byte]int)
	var maxLen, left = 0, 0
	for i := range s {
		var ch = s[i]
		if idx, ok := hash[ch]; ok && left <= idx {
			left = idx + 1
		}
		hash[ch] = i
		maxLen = max(maxLen, i-left+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
