package problem1101_1200

func tribonacci(n int) int {
	var a, b, c = 0, 1, 1
	for n > 0 {
		a, b, c = b, c, a+b+c
		n--
	}
	return a
}
