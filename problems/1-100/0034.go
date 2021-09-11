package problem1_100

func searchRange(nums []int, target int) []int {
	var low, high = 0, len(nums) - 1
	var ans = []int{-1, -1}
	// 找起始
	for low <= high {
		var mid = low + (high-low)>>1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				ans[0] = mid
				break
			}
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if ans[0] == -1 {
		return ans
	}
	// 找结束
	low, high = 0, len(nums)-1
	for low <= high {
		var mid = low + (high-low)>>1
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				ans[1] = mid
				break
			}
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
