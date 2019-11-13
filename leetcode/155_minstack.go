package main

/**
 * leetcode@155 最小栈
 * 要求常数时间检索到栈内最小元素
 */

type Node struct {
	next *Node
	val  int
}

type MinStack struct {
	dataTop *Node
	auxTop  *Node // 辅助栈
}

// leetcode 给的是大写,因为冲突改成小写,复制到 leetcode 需要记得改一下
func constructor() MinStack {
	return MinStack{nil, nil}
}

func (m *MinStack) Push(x int) {
	if m.auxTop == nil || x <= m.auxTop.val {
		m.auxTop = &Node{next: m.auxTop, val: x}
	}
	m.dataTop = &Node{next: m.dataTop, val: x}
}

func (m *MinStack) Pop() {
	if m.dataTop != nil {
		if m.dataTop.val == m.auxTop.val {
			m.auxTop = m.auxTop.next
		}
		m.dataTop = m.dataTop.next
	}
}

func (m *MinStack) Top() int {
	if m.dataTop != nil {
		return m.dataTop.val
	}
	return 0
}

func (m *MinStack) GetMin() int {
	if m.auxTop != nil {
		return m.auxTop.val
	}
	return 0
}
