package main

type StackNode struct {
	next *StackNode
	val  int
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return make([]int, 0)
	}

	var top *StackNode
	var hash = make(map[int]int)

	top = &StackNode{next: nil, val: nums2[0]}
	for i := 1; i < len(nums2); i++ {
		var now = nums2[i]

		if now < top.val {
			top = &StackNode{next: top, val: now}
		} else {
			for top != nil && top.val < now {
				hash[top.val] = now
				top = top.next
			}
			top = &StackNode{next: top, val: now}
		}
	}

	var result = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := hash[nums1[i]]; ok {
			result[i] = v
		} else {
			result[i] = -1
		}
	}
	return result
}
