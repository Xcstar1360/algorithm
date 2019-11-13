package main

import "fmt"

/**
 * 循环队列
 * 主要解决数组实现队列数据搬运问题
 * 会浪费一个存储空间
 */
type CircularQueue struct {
	items []interface{}
	n     int
	head  int
	tail  int
}

func NewCircularQueue(capacity int) CircularQueue {
	return CircularQueue{
		items: make([]interface{}, capacity),
		n:     capacity,
		head:  0,
		tail:  0,
	}
}

func (cq *CircularQueue) Enqueue(v interface{}) bool {
	if (cq.tail+1)%cq.n == cq.head {
		return false
	}

	cq.items[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.n
	return true
}

func (cq *CircularQueue) Dequeue() interface{} {
	if cq.head == cq.tail {
		return nil
	}

	var ret = cq.items[cq.head]
	cq.head = (cq.head + 1) % cq.n
	return ret
}

func (cq *CircularQueue) String() string {
	if cq.head == cq.tail {
		return "empty queue"
	}

	var i = cq.head
	var format = "head"
	for i != cq.tail {
		format += fmt.Sprintf(" <- %v", cq.items[i])
		i = (i + 1) % cq.n
	}
	format += " <- tail"
	return format
}
