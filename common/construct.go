package common

func ConstructList(nums []int) *ListNode {
	var p = &ListNode{}
	var l = p
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return l.Next
}

func CalcPrimeNumberQueue(N int) []int {
	var primes = make([]int, N)
	var pc, m, k int

	// 构造素数序列
	primes[0] = 2
	// 计数器
	pc = 1
	// 从3开始寻找
	m = 3
	for pc < N {
		k = 0
		// m若不为素数，必有一个素数因子小于m，则这个因子一定在primes数组序列中
		for primes[k]*primes[k] <= m {
			// 找到一个素数，则m加2，跳过偶数
			if m%primes[k] == 0 {
				m += 2
				k = 1
			} else {
				k++
			}
		}
		// 找不到素数因子，则说明m是素数，添加入素数序列，然后m加2，继续寻找
		primes[pc], pc = m, pc+1
		m += 2
	}

	return primes
}
