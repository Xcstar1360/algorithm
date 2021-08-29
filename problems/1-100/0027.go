package problem1_100

func removeElement(nums []int, val int) int {
	if len(nums) == 0 || val >= 51 {
		return len(nums)
	}

	var i, j = 0, len(nums) - 1
	for i <= j {
		if nums[i] == val {
			for nums[j] == val && j > i {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}

	return i
}
