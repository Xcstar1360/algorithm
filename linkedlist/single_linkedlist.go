package linkedlist

import "fmt"

/**
 * 单链表基础操作
 */

type ListNode struct {
	next 	*ListNode
	value 	interface{}	// 可以存储任意类型
}

type LinkedList struct {
	head 	*ListNode
	length 	int
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func (l *ListNode) GetValue() interface{} {
	return l.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

/**
 * 在某个节点之后插入节点
 */
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = p.next
	p.next = newNode
	l.length++
	return true
}

/**
 * 在某个节点之前插入节点
 */
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head
	// 查找 p 节点的前一个节点
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = cur
	pre.next = newNode
	l.length++
	return true
}

/**
 * 链表头部插入节点
 */
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

/**
 * 链表尾部插入节点
 */
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}

/**
 * 通过索引查找节点
 */
func (l *LinkedList) FindByIndex(index int) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

/**
 * 删除传入的节点
 */
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}

	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}
	pre.next = p.next
	p = nil	// GC
	l.length--
	return true
}

/**
 * 打印链表
 */
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
