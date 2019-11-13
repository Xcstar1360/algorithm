package main

import "fmt"

type ArrayQueue struct {
	items []interface{}
	n     int
	head  int
	tail  int
}

func NewArrayQueue(capacity int) ArrayQueue {
	return ArrayQueue{
		items: make([]interface{}, capacity),
		n:     capacity,
		head:  0,
		tail:  0,
	}
}

/**
 * 入队
 * 摊还分析时间复杂度为 O(1)
 */
func (aq *ArrayQueue) Enqueue(item interface{}) bool {
	if aq.tail == aq.n {
		if aq.head == 0 {
			return false
		}
		// 搬运
		for i := aq.head; i < aq.tail; i++ {
			aq.items[i-aq.head] = aq.items[i]
		}
		aq.tail -= aq.head
		aq.head = 0
	}
	aq.items[aq.tail] = item
	aq.tail++
	return true
}

func (aq *ArrayQueue) Dequeue() interface{} {
	if aq.head == aq.tail {
		return nil
	}

	var ret = aq.items[aq.head]
	aq.head++
	return ret
}

func (aq *ArrayQueue) String() string {
	if aq.head == aq.tail {
		return "empty queue"
	}
	var format = "head"
	for i := aq.head; i < aq.tail; i++ {
		format += fmt.Sprintf(" <- %v", aq.items[i])
	}
	format += " <- tail"
	return format
}
