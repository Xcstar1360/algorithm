package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 4),
		top:  -1,
	}
}

func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}

	if as.top > len(as.data)-1 {
		// 扩容
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	var val = as.data[as.top]
	as.top--
	return val
}

func (as *ArrayStack) IsEmpty() bool {
	return as.top < 0
}

func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

func (as *ArrayStack) Flush() {
	as.top = -1
}

func (as *ArrayStack) Print() string {
	if as.IsEmpty() {
		return "empty stack"
	}
	var format = "*"
	for i := as.top; i >= 0; i-- {
		format += fmt.Sprintf("%v", as.data[i])
		if i != 0 {
			format += "->"
		}
	}
	return format
}
