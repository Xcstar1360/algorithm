package stack

import "fmt"

type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (ls *LinkedListStack) Push(v interface{}) {
	ls.topNode = &node{next: ls.topNode, val: v}
}

func (ls *LinkedListStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	var val = ls.topNode.val
	ls.topNode = ls.topNode.next
	return val
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.topNode == nil
}

func (ls *LinkedListStack) Top() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	return ls.topNode.val
}

func (ls *LinkedListStack) Flush() {
	ls.topNode = nil
}

func (ls *LinkedListStack) Print() string {
	if ls.IsEmpty() {
		return "empty stack"
	}
	var format = "*"
	var cur = ls.topNode
	for cur != nil {
		format += fmt.Sprintf("%v", cur.val)
		if cur.next != nil {
			format += "->"
		}
		cur = cur.next
	}
	return format
}
