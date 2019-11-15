package main

import "fmt"

/**
 * 经典排序算法链表实现
 */

func main() {
	var head *ListNode
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	for _, v := range a {
		head = &ListNode{Val: v, Next: head}
	}
	printLinkedList(head)
	head = mergeSort(head)
	printLinkedList(head)
}

func printLinkedList(head *ListNode) {
	var cur, format = head, ""
	for cur != nil {
		format += fmt.Sprintf("%v -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println(format)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return mergeSortInternally(head)
}

func mergeSortInternally(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var preSlow *ListNode
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	preSlow.Next = nil
	var left = mergeSortInternally(head)
	var right = mergeSortInternally(slow)

	return merge(left, right)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead = &ListNode{Val: 0, Next: nil}
	var cur = dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummyHead.Next
}
