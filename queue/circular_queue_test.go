package main

import "testing"

func TestCircularQueue_Enqueue(t *testing.T) {
	var cq = NewCircularQueue(5)
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	cq.Enqueue(4)
	cq.Enqueue(5)
	cq.Enqueue(6)
	t.Log(cq.String())
}

func TestCircularQueue_Dequeue(t *testing.T) {
	var cq = NewCircularQueue(5)
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	cq.Enqueue(4)
	cq.Enqueue(5)
	cq.Enqueue(6)
	t.Log(cq.String())
	t.Log(cq.Dequeue())
	cq.Enqueue(6)
	t.Log(cq.String())
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	t.Log(cq.String())
}
