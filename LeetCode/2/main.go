package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode() *ListNode {
	return &ListNode{Val: 0, Next: nil}
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	nextTen := 0
	startNode := &ListNode{}
	resultNode := startNode

	for nextTen != 0 || l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := (v1 + v2 + nextTen)
		nextTen = sum / 10
		resultNode.Next = NewNode()
		resultNode.Next.Val = sum % 10
		resultNode = resultNode.Next

	}
	return startNode.Next
}

func main() {
	l1 := &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 9, Next: &ListNode{Val: 6, Next: nil}}}

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}

//      0 -> 2 -> 3
//      3 -> 9 -> 6
//      3 -> 1 -> 0 -> 1
