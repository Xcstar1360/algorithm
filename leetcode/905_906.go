package main

import "fmt"

func main() {
	var A = []int{3, 1, 2, 4}
	A = sortArrayByParity(A)
	fmt.Println(A)
	var B = []int{4, 2, 5, 7}
	B = sortArrayByParityII(B)
	fmt.Println(B)
}

func sortArrayByParity(A []int) []int {
	var j = 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}
	return A
}

func sortArrayByParityII(A []int) []int {
	var res = make([]int, len(A))
	var even, odd = 0, 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[even] = A[i]
			even += 2
		} else {
			res[odd] = A[i]
			odd += 2
		}
	}
	return res
}
