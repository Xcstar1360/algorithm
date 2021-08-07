package common

import (
	"bytes"

	"fmt"
)

/**
 * leetcode 题目给定的一些常用数据结构定义
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var buffer bytes.Buffer
	var len = 0
	for l != nil {
		buffer.WriteString(fmt.Sprint(l.Val))
		buffer.WriteString("->")
		len++
		l = l.Next
	}
	buffer.WriteString("nil")

	return fmt.Sprintf("List{size: %d, value: %s}", len, buffer.String())
}
