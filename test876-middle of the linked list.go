package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	fmt.Println(middleNode(&head))
}

func middleNode(head *ListNode) *ListNode {
	tmp := head
	length := 1
	for head.Next != nil {
		length++
		head = head.Next
	}
	middle := (length / 2) + 1

	count := 1
	for tmp.Next != nil {
		count++
		if count > middle {
			break
		}
		tmp = tmp.Next
	}

	return tmp
}
