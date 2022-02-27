package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l1Val  = 0
		l2Val  = 0
		l1Next = l1
		l2Next = l2
		sum    = 0
		head   = &ListNode{}
		curr   = head
	)
	for l1Next != nil || l2Next != nil {
		if l1Next != nil {
			l1Val = l1Next.Val
			l1Next = l1Next.Next
		} else {
			l1Val = 0
		}
		if l2Next != nil {
			l2Val = l2Next.Val
			l2Next = l2Next.Next
		} else {
			l2Val = 0
		}
		sum = l1Val + l2Val + sum
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		sum = sum / 10
	}
	if sum > 0 {
		curr.Next = &ListNode{Val: sum}
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	h := addTwoNumbers(l1, l2)
	fmt.Println(h.Val, h.Next.Val, h.Next.Next.Val)
}
