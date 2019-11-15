package main

import "fmt"

func main() {
	//var A = []int{1, 2, 3, 4, 5}
	var A = []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(A), A)
}

func pancakeSort(A []int) []int {
	var n = len(A)
	var res = make([]int, 0)
	for i := 0; i < n; i++ {
		var imax = 0
		var max = A[0]
		for j := 0; j < n-i; j++ {
			if A[j] > max {
				imax = j
				max = A[j]
			}
		}
		if imax != n-i-1 && imax != 0 {
			reverse(A, imax+1)
			res = append(res, imax+1)
			reverse(A, n-i)
			res = append(res, n-i)
		} else if imax == n-i-1 {
			// 啥也不做
		} else if imax == 0 {
			reverse(A, n-i)
			res = append(res, n-i)
		}
	}
	return res
}

// 翻转数组前 n 个元素
func reverse(A []int, n int) {
	var mid = n / 2
	for i := 0; i < mid; i++ {
		A[i], A[n-i-1] = A[n-i-1], A[i]
	}
}
