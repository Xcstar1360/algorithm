package main

import "fmt"

type Node struct {
	next *Node
	val  interface{}
}

type LinkedListQueue struct {
	head *Node
	tail *Node
}

func NewLinkedListQueue() LinkedListQueue {
	return LinkedListQueue{}
}

func (lq *LinkedListQueue) Enqueue(v interface{}) {
	if lq.tail == nil {
		var newNode = &Node{next: nil, val: v}
		lq.head = newNode
		lq.tail = newNode
	} else {
		lq.tail.next = &Node{next: nil, val: v}
		lq.tail = lq.tail.next
	}
}

func (lq *LinkedListQueue) Dequeue() interface{} {
	if lq.head == nil {
		return nil
	}

	var ret = lq.head.val
	lq.head = lq.head.next
	if lq.head == nil {
		lq.tail = nil
	}
	return ret
}

func (lq *LinkedListQueue) String() string {
	if lq.head == nil {
		return "empty queue"
	}

	var format = "head"
	var cur = lq.head
	for cur != nil {
		format += fmt.Sprintf(" <- %v", cur.val)
		cur = cur.next
	}
	format += " <- tail"
	return format
}
