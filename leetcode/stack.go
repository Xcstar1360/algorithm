package main

/**
 * 栈题目
 * leetcode@20 有效的括号(栈的经典应用)
 * leetcode@844 比较含退格的字符
 * leetcode@682 棒球比赛
 */

type StackNode struct {
	next *StackNode
	val  rune
}

/**
 * 字符串只包含 '(', ')', '[', ']', '{', '}'
 * 空字符串认为是有效字符串
 */
func isValid(s string) bool {
	if s == "" {
		return true
	}

	var leftBracket = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
	}
	var rightBracket = map[rune]int{
		')': -1,
		']': -2,
		'}': -3,
	}
	var top *StackNode
	for _, ch := range s {
		if _, ok := leftBracket[ch]; ok {
			top = &StackNode{next: top, val: ch}
		} else if rv, ok := rightBracket[ch]; ok {
			if top != nil {
				var lv = leftBracket[top.val]
				if lv+rv == 0 {
					top = top.next
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return top == nil
}

func strToStack(str string) *StackNode {
	var top *StackNode
	for _, ch := range str {
		if ch != '#' {
			top = &StackNode{next: top, val: ch}
		} else {
			if top != nil {
				top = top.next
			}
		}
	}
	return top
}

/**
 * 比较含退格的字符
 */
func backspaceCompare(S string, T string) bool {
	var top1 = strToStack(S)
	var top2 = strToStack(T)
	for top1 != nil && top2 != nil {
		if top1.val != top2.val {
			return false
		}
		top1 = top1.next
		top2 = top2.next
	}
	return top1 == nil && top2 == nil
}
