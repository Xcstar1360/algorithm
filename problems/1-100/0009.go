package problem1_100

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var tmp, rx = x, 0
	for x > 0 {
		rx = rx*10 + x%10
		x /= 10
	}
	return rx == tmp
}
