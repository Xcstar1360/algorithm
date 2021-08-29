package problem1_100

func removeDuplicates(nums []int) int {
	var n = len(nums)
	if n <= 1 {
		return n
	}

	var k = 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}
