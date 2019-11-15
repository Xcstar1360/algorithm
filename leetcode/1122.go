package main

/**
 * arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000 因为这个条件,所以直接使用 hash
 */
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res = make([]int, len(arr1))
	var hash = make([]int, 1001)
	for i := 0; i < len(arr1); i++ {
		hash[arr1[i]]++
	}
	var idx = 0
	for i := 0; i < len(arr2); i++ {
		for hash[arr2[i]] > 0 {
			res[idx] = arr2[i]
			idx++
			hash[arr2[i]]--
		}
	}
	for i := 0; i < len(hash); i++ {
		for hash[i] > 0 {
			res[idx] = i
			idx++
			hash[i]--
		}
	}
	return res
}
