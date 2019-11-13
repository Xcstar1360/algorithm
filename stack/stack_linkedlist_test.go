package stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	var as = NewLinedListStack()
	as.Push(1)
	as.Push(2)
	as.Push(4)
	t.Log(as.Print())
}

func TestLinkedListStack_Pop(t *testing.T) {
	var as = NewLinedListStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	as.Push(4)
	as.Push(5)
	t.Log(as.Print())
	as.Pop()
	as.Pop()
	as.Pop()
	as.Pop()
	t.Log(as.Print())
}

func TestLinkedListStack_Top(t *testing.T) {
	var as = NewLinedListStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	t.Log(as.Top())
	as.Pop()
	t.Log(as.Top())
}

func TestLinkedListStack_Flush(t *testing.T) {
	var as = NewLinedListStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	t.Log(as.Print())
	as.Flush()
	t.Log(as.Print())
}
