package one

/**
 * 解法一: 利用 hash
 */
func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i, v := range nums {
		if index, ok := hash[target-v]; ok {
			return []int{index, i}
		}
		hash[v] = i
	}
	return []int{}
}

/**
 * 解法二: 双循环暴力
 */
// func twoSum(nums []int, target int) []int {
// 	var n = len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
