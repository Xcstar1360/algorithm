package main

import "fmt"

func main() {
	var nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}

func findKthLargest(nums []int, k int) int {
	var n = len(nums)
	var p, r = 0, n - 1
	var q = partition(nums, p, r)
	k = n - k + 1

	for q+1 != k {
		if q+1 < k {
			p = q + 1
		} else {
			r = q - 1
		}
		q = partition(nums, p, r)
	}
	return nums[q]
}

func partition(a []int, p int, r int) int {
	var pivot = a[r]
	var i = p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}
