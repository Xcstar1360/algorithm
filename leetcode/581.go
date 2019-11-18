package main

import "fmt"

func main() {
	var arr = []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(arr))
	arr = []int{2, 2, 2, 3, 2}
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(nums []int) int {
	var n = len(nums)
	if n <= 1 {
		return 0
	}

	var max, high = nums[0], 0
	var min, low = nums[n-1], n - 1
	for i := 0; i < n; i++ {
		if nums[i] < max {
			high = i
		}
		if nums[i] > max {
			max = nums[i]
		}

		if nums[n-1-i] > min {
			low = n - 1 - i
		}
		if nums[n-1-i] < min {
			min = nums[n-1-i]
		}
	}

	if high > low {
		return high - low + 1
	}
	return 0
}
