package main

func removeDuplicates(nums []int) int {
	var n = len(nums)
	if n <= 1 {
		return n
	}

	var i, j = 0, 1
	for j < n {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}
