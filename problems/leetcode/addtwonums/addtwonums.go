package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode
	var l2 *ListNode

	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)
	l1 = lappend(l1, 9)

	l2 = lappend(l2, 9)
	l2 = lappend(l2, 9)
	l2 = lappend(l2, 9)
	l2 = lappend(l2, 9)

	lprint(l1)
	lprint(l2)

	l3 := addTwoNumbers(l1, l2)
	lprint(l3)
}

//go:inline
func lprint(head *ListNode) {
	if head == nil {
		fmt.Println("<nil>")
		return
	}
	for h := head; h != nil; h = h.Next {
		fmt.Printf("| %v | -> ", h.Val)
	}
	fmt.Println()
}

//go:inline
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//go:inline
func lappend(head *ListNode, v int) *ListNode {

	if head == nil {
		head = &ListNode{
			Val:  v,
			Next: nil,
		}
		return head
	}
	tmp := &ListNode{
		Val:  v,
		Next: nil,
	}
	l := head
	for ; l.Next != nil; l = l.Next {
	}
	l.Next = tmp
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head *ListNode

	carry := 0
	i := l1
	j := l2
	for {
		if i == nil || j == nil {
			break
		}
		val := (carry + i.Val + j.Val) % 10
		carry = (carry + i.Val + j.Val) / 10
		head = lappend(head, val)
		i = i.Next
		j = j.Next
	}
	remaining := i
	if i == nil {
		remaining = j
	}
	for ; remaining != nil; remaining = remaining.Next {
		val := (carry + remaining.Val) % 10
		carry = (carry + remaining.Val) / 10
		head = lappend(head, val)
	}
	if carry > 0 {
		head = lappend(head, carry)
	}
	return head
}
