package problem1_100

type StackNode struct {
	next *StackNode
	val  byte
}

var brackets = map[byte]int{
	'(': -1,
	')': 1,
	'[': -2,
	']': 2,
	'{': -3,
	'}': 3,
}

func isValid(s string) bool {
	var stack *StackNode
	for i := range s {
		var v = brackets[s[i]]
		if v < 0 {
			stack = &StackNode{val: s[i], next: stack}
		} else {
			if stack == nil {
				return false
			}
			var v1 = brackets[stack.val]
			if v+v1 != 0 {
				return false
			}
			stack = stack.next
		}
	}
	return stack == nil
}
