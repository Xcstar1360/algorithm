package main

/**
 * leetcode@224 基本计算器
 * 只有 ( ) + - 非负整数与空格,并且可以假设给定表达式都有效
 */

var levels = map[uint8]int{
	'+': 1,
	'-': 1,
	')': 0,
	'(': 2,
	' ': -1,
}

// 没有泛型,只能定义两个
type operandNode struct {
	next *operandNode
	val  int
}

type operatorNode struct {
	next *operatorNode
	val  uint8
}

func calculate(s string) int {
	return 0
}
