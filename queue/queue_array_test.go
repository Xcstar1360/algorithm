package main

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
	var aq = NewArrayQueue(5)
	aq.Enqueue(1)
	aq.Enqueue(2)
	aq.Enqueue(3)
	aq.Enqueue(4)
	aq.Enqueue(5)
	aq.Enqueue(6)
	t.Log(aq.String())
}

func TestArrayQueue_Dequeue(t *testing.T) {
	var aq = NewArrayQueue(5)
	aq.Enqueue(1)
	aq.Enqueue(2)
	aq.Enqueue(3)
	aq.Enqueue(4)
	aq.Enqueue(5)
	aq.Enqueue(6)
	t.Log(aq.String())
	t.Log(aq.Dequeue())
	aq.Enqueue(6)
	t.Log(aq.String())
	aq.Dequeue()
	aq.Dequeue()
	aq.Dequeue()
	aq.Dequeue()
	aq.Dequeue()
	t.Log(aq.String())
}
