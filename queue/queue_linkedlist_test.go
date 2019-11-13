package main

import "testing"

func TestLinkedListQueue_Enqueue(t *testing.T) {
	var lq = NewLinkedListQueue()
	lq.Enqueue(1)
	lq.Enqueue(2)
	lq.Enqueue(3)
	lq.Enqueue(4)
	lq.Enqueue(5)
	lq.Enqueue(6)
	t.Log(lq.String())
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	var lq = NewLinkedListQueue()
	lq.Enqueue(1)
	lq.Enqueue(2)
	lq.Enqueue(3)
	lq.Enqueue(4)
	lq.Enqueue(5)
	lq.Enqueue(6)
	t.Log(lq.String())
	t.Log(lq.Dequeue())
	lq.Enqueue(6)
	t.Log(lq.String())
	lq.Dequeue()
	lq.Dequeue()
	lq.Dequeue()
	lq.Dequeue()
	lq.Dequeue()
	t.Log(lq.String())
}
