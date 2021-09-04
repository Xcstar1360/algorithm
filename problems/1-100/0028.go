package problem1_100

/**
 * 0 <= haystack.length, needle.length <= 5 * 10^4
 * haystack 和 needle 仅由小写英文字符组成
 */

/**
 * BF 暴力
 * 执行用时: 208ms,内存消耗: 2.6M
 */
func strStrBF(haystack string, needle string) int {
	var n, m = len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	if n == 0 || m > n {
		return -1
	}
	for i := 0; i < n-m+1; i++ {
		var j = 0
		for j < m {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}

/**
 * 快速生成素数序列：
 * github.com/Xcstar1360/algorithm/common/construct.go
 * func CalcPrimeNumberQueue(N int) []int
 */
var primeMap = map[byte]int{
	'a': 2,
	'b': 3,
	'c': 5,
	'd': 7,
	'e': 11,
	'f': 13,
	'g': 17,
	'h': 19,
	'i': 23,
	'j': 29,
	'k': 31,
	'l': 37,
	'm': 41,
	'n': 43,
	'o': 47,
	'p': 53,
	'q': 59,
	'r': 61,
	's': 67,
	't': 71,
	'u': 73,
	'v': 79,
	'w': 83,
	'x': 89,
	'y': 97,
	'z': 101,
}

// s 仅由小写字母组成
func calcHashCode(s string) int {
	var hashCode = 0
	for i := 0; i < len(s); i++ {
		hashCode += primeMap[s[i]]
	}
	return hashCode
}

/**
 * RK
 * 执行用时: 4ms,内存消耗: 2.9M
 * 时间复杂度很大程度取决于 hash 算法的设计
 * 1. 子串长度范围很大，不适用 K 进制 hash，PASS
 * 2. a-z 对应 1-26，很容易冲突，导致不必要的比较，提升时间复杂度，PASS
 * 3. 将 a-z 对应到素数，PICK
 */
func strStrRK(haystack string, needle string) int {
	var n, m = len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	if n == 0 || m > n {
		return -1
	}

	var needleHashCode = calcHashCode(needle)
	// 预计算，利用一定的优化避免双循环计算 hash
	var hashCodes = make([]int, n-m+1)
	hashCodes[0] = calcHashCode(haystack[0:m])
	for i := 1; i < n-m+1; i++ {
		hashCodes[i] = hashCodes[i-1] -
			primeMap[haystack[i-1]] +
			primeMap[haystack[i+m-1]]
	}
	for i := 0; i < n-m+1; i++ {
		if hashCodes[i] == needleHashCode {
			var j = 0
			for j < m {
				if haystack[i+j] != needle[j] {
					break
				}
				j++
			}
			if j == m {
				return i
			}
		}
	}
	return -1
}

/**
 * BM 与 KMP
 */
