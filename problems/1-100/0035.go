package problem1_100

func searchInsert(nums []int, target int) int {
	var low, high = 0, len(nums) - 1
	for low <= high {
		var mid = low + (high-low)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return len(nums)
}
