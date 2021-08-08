package problem1_100

import "bytes"

/**
 * 题目规定只有英文字母大小写以及 "." ","
 * 就不考虑使用 rune 了
 */
func convert(s string, numRows int) string {
	var k = len(s)
	if k <= numRows || numRows == 1 {
		return s
	}

	var n, step = numRows, 2*numRows - 2
	var buffer bytes.Buffer
	for n > 0 {
		var next, start = 2*n - 2, numRows - n
		for start < k {
			buffer.WriteByte(s[start])
			if n != numRows && n != 1 && start+next < k {
				buffer.WriteByte(s[start+next])
			}
			start += step
		}
		n--
	}
	return buffer.String()
}
